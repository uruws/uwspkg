// Copyright (c) Jeremías Casteglione <jeremias@talkingpts.org>
// See LICENSE file.

// Package build implements build steps.
package build

import (
	"uwspkg/log"
	"uwspkg/manifest"
)

func Package(m *manifest.Config) error {
	log.Debug("package %s", m.Origin)
	return nil
}
