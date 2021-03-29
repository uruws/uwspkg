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
		1: cfg.DebianInstallVariant,
		2: cfg.DebianRepo,
		3: cfg.DebianSecRepo,
		4: dist,
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
	log.Print("SetUp %s build.", m.Origin)
	sess := "uwspkg-build-" + m.Session
	if err := profile.Create(cfg, m, sess); err != nil {
		return err
	}
	err := libexec.Run("build/session-start", sess)
	if err != nil {
		if err := profile.Remove(cfg, m, sess); err != nil {
			log.Error("%v", err)
		}
		return err
	}
	m.SessionStart = time.Now()
	return nil
}

func TearDown(cfg *config.Main, m *manifest.Config) []error {
	log.Print("TearDown %s build.", m.Origin)
	errlist := make([]error, 0)
	if m.SessionStart.IsZero() {
		return errlist
	}
	sess := "uwspkg-build-" + m.Session
	if err := libexec.Run("build/session-stop", sess); err != nil {
		errlist = append(errlist, err)
	} else {
		if err := profile.Remove(cfg, m, sess); err != nil {
			errlist = append(errlist, err)
		}
	}
	return errlist
}

func Source(m *manifest.Config) error {
	var err error
	log.Debug("make source %s.", m.Origin)
	sess := "uwspkg-build-" + m.Session
	err = libexec.Run("build/make-fetch", sess, m.Origin, m.Name, m.Fetch)
	if err != nil {
		return err
	}
	err = libexec.Run("build/source-package", sess, m.Origin, m.Name, m.Fetch)
	if err != nil {
		return err
	}
	return nil
}

func Package(m *manifest.Config) error {
	log.Print("Make %s.", m.Origin)
	return nil
}
