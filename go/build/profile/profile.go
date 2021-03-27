// Copyright (c) Jeremías Casteglione <jeremias@talkingpts.org>
// See LICENSE file.

// Package profile implements a build profiles manager.
package profile

import (
	"uwspkg/config"
	"uwspkg/log"
	"uwspkg/manifest"
)

func SetUp(cfg *config.Main, m *manifest.Config) error {
	log.Debug("setup %s %s", m.Origin, m.Profile)
	return nil
}
