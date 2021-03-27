// Copyright (c) Jeremías Casteglione <jeremias@talkingpts.org>
// See LICENSE file.

// Package profile implements a build profiles manager.
package profile

import (
	"uwspkg/config"
	//~ "uwspkg/libexec"
	"uwspkg/log"
	"uwspkg/manifest"
)

func Create(cfg *config.Main, m *manifest.Config) error {
	log.Debug("%s create %s %s", m.Session, m.Origin, m.Profile)
	args := []string{
		0: m.Profile,
		1: "build-"+m.Session,
		2: cfg.SchrootCfgDir,
	}
	//~ return libexec.Run("build/profile-create", args...)
	log.Debug("args: %v", args)
	return nil
}

func Remove(cfg *config.Main, m *manifest.Config) error {
	log.Debug("%s remove %s %s", m.Session, m.Origin, m.Profile)
	args := []string{
		0: m.Profile,
		1: "build-"+m.Session,
		2: cfg.SchrootCfgDir,
	}
	//~ return libexec.Run("build/profile-remove", args...)
	log.Debug("args: %v", args)
	return nil
}
