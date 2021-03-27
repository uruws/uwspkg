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
	log.Info("SetUp %s build.", m.Origin)
	if err := profile.Create(cfg, m); err != nil {
		return err
	}
	err := libexec.Run("build/session-start", "build-"+m.Session, m.Profile)
	if err != nil {
		return err
	}
	m.SessionStart = time.Now()
	return nil
}

func TearDown(cfg *config.Main, m *manifest.Config) []error {
	log.Info("TearDown %s build.", m.Origin)
	errlist := make([]error, 0)
	if m.SessionStart.IsZero() {
		return errlist
	}
	if err := libexec.Run("build/session-stop", "build-"+m.Session); err != nil {
		errlist = append(errlist, err)
	} else {
		if err := profile.Remove(cfg, m); err != nil {
			errlist = append(errlist, err)
		}
	}
	return errlist
}

func Package(m *manifest.Config) error {
	log.Info("Make %s.", m.Origin)
	return nil
}
