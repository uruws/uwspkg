// Copyright (c) Jeremías Casteglione <jeremias@talkingpts.org>
// See LICENSE file.

// Package build implements build steps.
package build

import (
	"io/ioutil"
	"path/filepath"
	"strings"
	"time"

	"uwspkg/build/profile"
	"uwspkg/config"
	"uwspkg/libexec"
	"uwspkg/log"
	"uwspkg/manifest"
)

func Bootstrap(cfg *config.Main) error {
	log.Info("Bootstrap FreeBSD pkg")
	if err := buildSetup(cfg); err != nil {
		return log.DebugError(err)
	}
	if err := setupProfile(cfg, "clang"); err != nil {
		return log.DebugError(err)
	}
	distfn := filepath.Join(cfg.SchrootCfgDir, "uwspkg-clang", "debian.distro")
	distro := ""
	if blob, err := ioutil.ReadFile(distfn); err != nil {
		return log.DebugError(err)
	} else {
		distro = strings.TrimSpace(string(blob))
	}
	log.Debug("bootstrap debian distro: %s", distro)
	if err := debianInstall(cfg, distro); err != nil {
		return log.DebugError(err)
	}
	if err := debianInstallProfile(cfg, "clang"); err != nil {
		return log.DebugError(err)
	}
	return nil
}

func EnvSetUp(cfg *config.Main) error {
	log.Info("Setup build environment.")
	log.Debug("schroot config: %s -> %s", cfg.BuildCfgDir, cfg.SchrootCfgDir)
	log.Debug("build dir: %s", cfg.BuildDir)
	log.Debug("packages dir: %s", cfg.PkgDir)
	log.Debug("debian install: %s", cfg.DebianInstall)
	log.Info("Debian deps: %s.", strings.Join(cfg.DebianDeps, " "))
	if err := buildSetup(cfg); err != nil {
		return err
	}
	for _, prof := range cfg.BuildProfile {
		if prof == "default" {
			continue
		}
		if err := setupProfile(cfg, prof); err != nil {
			return err
		}
	}
	for _, dist := range cfg.DebianDistro {
		if err := debianInstall(cfg, dist); err != nil {
			return err
		}
	}
	if err := debianInstallProfile(cfg, "internal"); err != nil {
		return err
	}
	for _, prof := range cfg.BuildProfile {
		if prof == "default" {
			continue
		}
		if err := debianInstallProfile(cfg, prof); err != nil {
			return err
		}
	}
	return nil
}

func buildSetup(cfg *config.Main) error {
	args := []string{
		0: cfg.BuildDir,
		1: cfg.BuildCfgDir,
		2: cfg.SchrootCfgDir,
		3: cfg.PkgDir,
	}
	for _, pkg := range cfg.DebianDeps {
		args = append(args, pkg)
	}
	return libexec.Run("build/setup", args...)
}

func setupProfile(cfg *config.Main, prof string) error {
	log.Info("Setup profile: %s.", prof)
	args := []string{
		0: cfg.BuildDir,
		1: cfg.BuildCfgDir,
		2: cfg.SchrootCfgDir,
		3: prof,
	}
	return libexec.Run("build/setup-profile", args...)
}

func debianInstall(cfg *config.Main, dist string) error {
	log.Info("Debian install %s.", dist)
	args := []string{
		0: cfg.BuildDir,
		1: cfg.BuildCfgDir,
		2: cfg.DebianInstallVariant,
		3: cfg.DebianRepo,
		4: cfg.DebianSecRepo,
		5: dist,
	}
	return libexec.Run("build/debian-install", args...)
}

func debianInstallProfile(cfg *config.Main, prof string) error {
	log.Info("Debian install profile: %s.", prof)
	e := libexec.EnvConfig(cfg)
	return libexec.RunEnv(e, "build/debian-install-profile", prof)
}

func SetUp(cfg *config.Main, m *manifest.Config) error {
	log.Print("Build %s setup.", m.Package)
	m.SessionStart = time.Now()
	if err := profile.Create(cfg, m); err != nil {
		return err
	}
	return nil
}

func TearDown(m *manifest.Config) []error {
	log.Print("Build %s tear down.", m.Package)
	errlist := make([]error, 0)
	if err := profile.Remove(m); err != nil {
		errlist = append(errlist, err)
	}
	return errlist
}

func Source(m *manifest.Config) error {
	var err error
	log.Print("Build %s source archive.", m.Package)
	srcdir := filepath.FromSlash("/uwspkg/src")
	// fetch
	chroot := libexec.NewChroot(m.BuildSession)
	chroot.Dir(filepath.Join(srcdir, m.OriginPath))
	err = chroot.Run(m.Environ(), "internal/make-fetch", m.Fetch)
	if err != nil {
		return err
	}
	// archive
	chroot.Dir("/build")
	err = chroot.Run(m.Environ(), "internal/source-archive")
	if err != nil {
		return err
	}
	return nil
}

func Package(m *manifest.Config) error {
	var err error
	log.Print("Build %s...", m.Package)
	log.Debug("%s build: %s", m.Session, m.Package)
	srcdir := filepath.FromSlash("/uwspkg/src")
	// chroot session
	chroot := libexec.NewChroot(m.BuildSession)
	chroot.Dir(filepath.Join(srcdir, m.OriginPath))
	err = chroot.SessionBegin("build-sess-"+m.Session)
	if err != nil {
		return err
	}
	defer chroot.SessionEnd()
	// depends
	err = chroot.Run(m.Environ(), "internal/make", m.Depends)
	if err != nil {
		return err
	}
	// build
	err = chroot.Run(m.Environ(), "internal/make", m.Build)
	if err != nil {
		return err
	}
	// check
	err = chroot.Run(m.Environ(), "internal/make", m.Check)
	if err != nil {
		return err
	}
	// install
	err = chroot.Run(m.Environ(), "internal/make-install", m.Install)
	if err != nil {
		return err
	}
	chroot.SessionEnd()
	return buildPackage(m)
}

func buildPackage(m *manifest.Config) error {
	log.Debug("%s build package: %s", m.Session, m.Package)
	chroot := libexec.NewChroot("internal-uwspkg")
	chroot.Dir("/build")
	return chroot.Run(m.Environ(), "internal/make-package")
}
