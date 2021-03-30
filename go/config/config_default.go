// Copyright (c) Jeremías Casteglione <jeremias@talkingpts.org>
// See LICENSE file.

// +build !uwspkg_build_package

package config

import (
	"path/filepath"
)

func init() {
	libxd, err := filepath.Abs(filepath.FromSlash("./libexec/utils"))
	if err != nil {
		panic(err)
	}
	defaultConfig.Libexec = libxd
}
