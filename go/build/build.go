// Copyright (c) Jeremías Casteglione <jeremias@talkingpts.org>
// See LICENSE file.

// Package build implements build steps.
package build

import (
	"uwspkg/log"
	"uwspkg/manifest"
)

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
