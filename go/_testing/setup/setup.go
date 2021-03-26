// Copyright (c) Jeremías Casteglione <jeremias@talkingpts.org>
// See LICENSE file.

// Package setup implements testing setup.
package setup

import (
	"os"
	"path/filepath"

	"uwspkg/config"
	"uwspkg/log"
)

func init() {
	log.Init(filepath.Base(os.Args[0]))
	config.Files = map[int]string{
		0: filepath.FromSlash("/go/src/uwspkg/testdata/uwspkg.yml"),
		1: filepath.FromSlash("./testdata/uwspkg.yml"),
	}
}
