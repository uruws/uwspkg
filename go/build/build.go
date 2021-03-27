// Copyright (c) Jeremías Casteglione <jeremias@talkingpts.org>
// See LICENSE file.

// Package build implements build steps.
package build

import (
	"strings"

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
	log.Info("Debian deps: %s", strings.Join(cfg.DebianDeps, " "))
	if err := buildSetup(cfg); err != nil {
		return err
	}
	for _, prof := range cfg.BuildProfile {
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
	if err := libexec.Run("build/setup", args...); err != nil {
		return err
	}
	return nil
}

func setupProfile(cfg *config.Main, prof string) error {
	log.Info("Setup profile %s.", prof)
	args := []string{
		0: cfg.BuildDir,
		1: cfg.BuildCfgDir,
		2: cfg.SchrootCfgDir,
		3: prof,
	}
	if err := libexec.Run("build/setup-profile", args...); err != nil {
		return err
	}
	return nil
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
	if err := libexec.Run("build/debian-install", args...); err != nil {
		return err
	}
	return nil
}

func debianInstallProfile(cfg *config.Main, prof string) error {
	log.Info("Debian install profile %s.", prof)
	args := []string{
		0: cfg.BuildDir,
		1: cfg.BuildCfgDir,
		2: prof,
	}
	if err := libexec.Run("build/debian-install-profile", args...); err != nil {
		return err
	}
	return nil
}

func SetUp(m *manifest.Config) error {
	log.Info("SetUp %s", m.Origin)
	return nil
}

func Package(m *manifest.Config) error {
	log.Info("Build %s", m.Origin)
	return nil
}

func TearDown(m *manifest.Config) error {
	log.Info("TearDown %s", m.Origin)
	return nil
}
