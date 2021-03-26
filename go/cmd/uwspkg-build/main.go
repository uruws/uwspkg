// Copyright (c) Jeremías Casteglione <jeremias@talkingpts.org>
// See LICENSE file.

// Package main implements uwspkg-build cmd.
package main

import (
	"flag"
	"os"
	"path"

	"uwspkg"
	"uwspkg/config"
	"uwspkg/log"
)

func main() {
	log.Init("uwspkg-build")
	log.Debug("init")
	var (
		pkgorig string
		pkgdir  string
		pkgname string
	)
	flag.Parse()
	pkgorig = flag.Arg(0)
	pkgdir, pkgname = parseOrigin(pkgorig)
	if pkgdir == "" {
		usage()
	}
	log.Debug("pkg origin: %s - %s %s", pkgorig, pkgdir, pkgname)
	if err := config.Load(); err != nil {
		log.Fatal("%v", err)
	}
	pkg := uwspkg.New(pkgorig)
	if err := pkg.Load(); err != nil {
		log.Fatal("%v", err)
	}
}

func parseOrigin(o string) (string, string) {
	return path.Split(o)
}

func usage() {
	log.Error("no package origin")
	flag.PrintDefaults()
	os.Exit(1)
}
