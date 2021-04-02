// Copyright (c) Jeremías Casteglione <jeremias@talkingpts.org>
// See LICENSE file.

// +build !uwspkg_build_package

package config

import (
	fp "path/filepath"
)

func init() {
	if libxd, err := fp.Abs(fp.FromSlash("./libexec/utils")); err != nil {
		panic(err)
	} else {
		defaultConfig.Libexec = libxd
	}
	if cfgd, err := fp.Abs(fp.FromSlash("./etc/schroot")); err != nil {
		panic(err)
	} else {
		defaultConfig.BuildCfgDir = cfgd
	}
	if bstfn, err := fp.Abs(fp.FromSlash("../build/uwspkg-bootstrap.tgz")); err != nil {
		panic(err)
	} else {
		defaultConfig.PkgBootstrap = bstfn
	}
}
