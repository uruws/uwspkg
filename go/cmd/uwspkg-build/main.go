// Copyright (c) Jeremías Casteglione <jeremias@talkingpts.org>
// See LICENSE file.

// Package main implements uwspkg-build cmd.
package main

import (
	"flag"
	"os"
	"path"

	"uwspkg"
	"uwspkg/build"
	"uwspkg/config"
	"uwspkg/log"
)

func main() {
	log.Init("uwspkg-build")
	log.Debug("init")
	var (
		buildSetup bool
	)
	flag.BoolVar(&buildSetup, "setup", false, "setup build environment")
	flag.Parse()
	if buildSetup {
		if err := build.EnvSetUp(); err != nil {
			log.Fatal("%v", err)
		}
	} else {
		pkgBuild(flag.Arg(0))
	}
}

func pkgBuild(pkgorig string) {
	pkgdir, pkgname := parseOrigin(pkgorig)
	if pkgdir == "" {
		usage()
	}
	log.Debug("pkg origin: %s - %s %s", pkgorig, pkgdir, pkgname)
	var (
		cfg *config.Main
		err error
	)
	if cfg, err = config.Load(); err != nil {
		log.Fatal("%v", err)
	}
	pkg := uwspkg.New(pkgorig, cfg)
	if err := pkg.Load(); err != nil {
		log.Fatal("%v", err)
	}
	if err := pkg.Build(); err != nil {
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
