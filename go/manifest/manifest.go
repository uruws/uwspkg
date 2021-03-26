// Copyright (c) Jeremías Casteglione <jeremias@talkingpts.org>
// See LICENSE file.

// Package manifest defines the package manifest.
package manifest

type Config struct {
	Origin string
	Name   string
}

func New() *Config {
	return &Config{}
}

func Parse(m *Config) error {
	return nil
}
