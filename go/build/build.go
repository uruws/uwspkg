// Copyright (c) Jeremías Casteglione <jeremias@talkingpts.org>
// See LICENSE file.

// Package build implements build steps.
package build

import (
	"uwspkg/config"
	"uwspkg/libexec"
	"uwspkg/log"
	"uwspkg/manifest"
)

func EnvSetUp(cfg *config.Main) error {
	log.Info("Env setup: %s -> %s", cfg.BuildCfgDir, cfg.SchrootCfgDir)
	log.Debug("Build dir: %s", cfg.BuildDir)
	log.Debug("Debian install: %s", cfg.DebianInstall)
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
