// Copyright (c) Jeremías Casteglione <jeremias@talkingpts.org>
// See LICENSE file.

// Package build implements build steps.
package build

import (
	"uwspkg/config"
	"uwspkg/log"
	"uwspkg/manifest"
)

func EnvSetUp(cfg *config.Main) error {
	log.Info("Env setup: %s -> %s", cfg.BuildCfgDir, cfg.SchrootCfgDir)
	log.Debug("Build dir: %s", cfg.BuildDir)
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