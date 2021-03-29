// Copyright (c) Jeremías Casteglione <jeremias@talkingpts.org>
// See LICENSE file.

// Package build implements build steps.
package build

import (
	"strings"
	"time"

	"uwspkg/build/profile"
	"uwspkg/config"
	"uwspkg/libexec"
	"uwspkg/log"
	"uwspkg/manifest"
)

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
	args := []string{
		0: cfg.BuildDir,
		1: cfg.BuildCfgDir,
		2: prof,
	}
	return libexec.Run("build/debian-install-profile", args...)
}

func SetUp(cfg *config.Main, m *manifest.Config) error {
	log.Print("Build %s setup.", m.Package)
	m.SessionStart = time.Now()
	if err := profile.Create(m); err != nil {
		return err
	}
	return nil
}

func TearDown(cfg *config.Main, m *manifest.Config) []error {
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
	err = libexec.Run("build/make-fetch", m.BuildSession, m.Origin, m.Name, m.Fetch)
	if err != nil {
		return err
	}
	err = libexec.Run("build/source-archive", m.BuildSession, m.Package)
	if err != nil {
		return err
	}
	return nil
}

func Package(m *manifest.Config) error {
	var err error
	log.Print("Build %s.", m.Package)
	log.Debug("%s build: %s", m.Package, m.Build)
	err = libexec.Run("build/make", m.BuildSession, m.Origin, m.Build)
	if err != nil {
		return err
	}
	return nil
}
