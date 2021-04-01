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

func Create(cfg *config.Main, m *manifest.Config) error {
	log.Debug("%s create %s %s", m.Session, m.Origin, m.Profile)
	chroot := libexec.NewChroot("internal-uwspkg")
	chroot.User("root")
	return chroot.Run(m.Environ(), "internal/profile-create", cfg.BuildDir, cfg.PkgDir)
}

func Remove(m *manifest.Config) error {
	log.Debug("%s remove %s %s", m.Session, m.Origin, m.Profile)
	chroot := libexec.NewChroot("internal-uwspkg")
	chroot.User("root")
	return chroot.Run(m.Environ(), "internal/profile-remove")
}
