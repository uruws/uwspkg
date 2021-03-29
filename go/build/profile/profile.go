// Copyright (c) Jeremías Casteglione <jeremias@talkingpts.org>
// See LICENSE file.

// Package profile implements a build profiles manager.
package profile

import (
	"uwspkg/libexec"
	"uwspkg/log"
	"uwspkg/manifest"
)

func Create(m *manifest.Config) error {
	log.Debug("%s create %s %s", m.Session, m.Origin, m.Profile)
	return libexec.RunEnv(m.Environ(), "build/profile-create")
}

func Remove(m *manifest.Config) error {
	log.Debug("%s remove %s %s", m.Session, m.Origin, m.Profile)
	return libexec.RunEnv(m.Environ(), "build/profile-remove")
}
