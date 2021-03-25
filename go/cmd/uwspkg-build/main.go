// Copyright (c) Jeremías Casteglione <jeremias@talkingpts.org>
// See LICENSE file.

// Package main implements uwspkg-build cmd.
package main

import (
	"uwspkg/log"
)

func main() {
	log.Init("uwspkg-build")
	log.Debug("init")
}
