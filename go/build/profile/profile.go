// Copyright (c) Jeremías Casteglione <jeremias@talkingpts.org>
// See LICENSE file.

// Package profile implements a build profiles manager.
package profile

import (
	"uwspkg/config"
	"uwspkg/libexec"
	"uwspkg/log"
	"uwspkg/manifest"
)

func Create(cfg *config.Main, m *manifest.Config, sess string) error {
	log.Debug("%s create %s %s", m.Session, m.Origin, m.Profile)
	return libexec.Run("build/profile-create", m.Profile, sess, cfg.BuildDir)
}

func Remove(cfg *config.Main, m *manifest.Config, sess string) error {
	log.Debug("%s remove %s %s", m.Session, m.Origin, m.Profile)
	return libexec.Run("build/profile-remove", sess, cfg.BuildDir)
}
