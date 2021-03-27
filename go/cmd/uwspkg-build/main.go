// Copyright (c) Jeremías Casteglione <jeremias@talkingpts.org>
// See LICENSE file.

// Package main implements uwspkg-build cmd.
package main

import (
	"flag"
	"fmt"
	"os"
	"path"

	"uwspkg"
	"uwspkg/build"
	"uwspkg/config"
	"uwspkg/libexec"
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
	var (
		cfg *config.Main
		err error
	)
	if cfg, err = config.Load(); err != nil {
		log.Fatal("%v", err)
	}
	if err = libexec.Configure(cfg); err != nil {
		log.Fatal("%v", err)
	}
	if buildSetup {
		err = build.EnvSetUp(cfg)
	} else {
		err = pkgBuild(cfg, flag.Arg(0))
	}
	if err != nil {
		log.Fatal("%v", err)
	}
	log.Debug("end")
}

func pkgBuild(cfg *config.Main, pkgorig string) error {
	pkgdir, pkgname := parseOrigin(pkgorig)
	if pkgdir == "" {
		usage()
	}
	log.Debug("pkg origin: %s - %s %s", pkgorig, pkgdir, pkgname)
	pkg := uwspkg.New(pkgorig, cfg)
	if err := pkg.Load(); err != nil {
		return err
	}
	if err := pkg.Build(); err != nil {
		return err
	}
	return nil
}

func parseOrigin(o string) (string, string) {
	return path.Split(o)
}

func usage() {
	log.Error("no package origin")
	fmt.Fprintf(os.Stderr, "Usage: uwspkg-build pkg/origin\n")
	flag.PrintDefaults()
	os.Exit(1)
}
