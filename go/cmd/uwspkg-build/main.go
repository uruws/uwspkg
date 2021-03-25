// Copyright (c) Jeremías Casteglione <jeremias@talkingpts.org>
// See LICENSE file.

// Package main implements uwspkg-build cmd.
package main

import (
	"flag"
	"os"
	"path"

	"uwspkg/log"
)

func main() {
	log.Init("uwspkg-build")
	log.Debug("init")
	var (
		pkgdir string
		pkgname string
	)
	flag.Parse()
	pkgdir, pkgname = parseOrigin(flag.Arg(0))
	if pkgdir == "" {
		usage()
	}
	log.Debug("pkg origin: %s%s", pkgdir, pkgname)
}

func parseOrigin(o string) (string, string) {
	return path.Split(o)
}

func usage() {
	log.Error("no package origin")
	flag.PrintDefaults()
	os.Exit(1)
}
