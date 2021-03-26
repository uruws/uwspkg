// Copyright (c) Jeremías Casteglione <jeremias@talkingpts.org>
// See LICENSE file.

// Package build implements build steps.
package build

import (
	"uwspkg/log"
	"uwspkg/manifest"
)

func SetUp(m *manifest.Config) error {
	log.Info("setup %s", m.Origin)
	return nil
}

func Package(m *manifest.Config) error {
	log.Info("package %s", m.Origin)
	return nil
}

func TearDown(m *manifest.Config) error {
	log.Info("tear down %s", m.Origin)
	return nil
}
