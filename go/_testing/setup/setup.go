// Copyright (c) Jeremías Casteglione <jeremias@talkingpts.org>
// See LICENSE file.

// Package setup implements testing setup.
package setup

import (
	"path/filepath"

	"uwspkg/config"
	"uwspkg/log"
)

func init() {
	log.Init("testing")
	config.ConfigFiles = map[int]string{
		0: filepath.FromSlash("./testdata"),
	}
}
