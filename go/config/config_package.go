// Copyright (c) Jeremías Casteglione <jeremias@talkingpts.org>
// See LICENSE file.

// +build uwspkg_build_package

package config

import (
	"path/filepath"
)

func init() {
	defaultConfig.Libexec = filepath.FromSlash("/uws/libexec/uwspkg")
	defaultConfig.BuildCfgDir = filepath.FromSlash("/uws/etc/schroot")
	defaultConfig.PkgBootstrap = filepath.FromSlash("/uws/lib/uwspkg/bootstrap.tgz")
}
