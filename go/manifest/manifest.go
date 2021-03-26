// Copyright (c) Jeremías Casteglione <jeremias@talkingpts.org>
// See LICENSE file.

// Package manifest defines the package manifest.
package manifest

import (
	"uwspkg/log"
)

type Config struct {
	Origin string
	Name   string
}

func newConfig() *Config {
	return &Config{}
}

type Manifest struct {
	cfg *Config
}

func New() *Manifest {
	return &Manifest{
		cfg: newConfig(),
	}
}

func (m *Manifest) Load(filename string) error {
	log.Debug("load %s", filename)
	return nil
}
