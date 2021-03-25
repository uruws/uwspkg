// Copyright (c) Jeremías Casteglione <jeremias@talkingpts.org>
// See LICENSE file.

// Package main implements uwspkg-build cmd.
package main

import (
	"flag"
	"os"

	"uwspkg/log"
)

func main() {
	log.Init("uwspkg-build")
	log.Debug("init")
	var (
		pkgorig string
	)
	flag.Parse()
	pkgorig = flag.Arg(0)
	if pkgorig == "" {
		usage()
	}
}

func usage() {
	log.Error("no package origin")
	flag.PrintDefaults()
	os.Exit(1)
}
