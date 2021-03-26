// Copyright (c) Jeremías Casteglione <jeremias@talkingpts.org>
// See LICENSE file.

// Package manifest defines the package manifest.
package manifest

type Config struct {
	Origin string
	Name   string
}

func New(origin, pkgname, filename string) (*Config, error) {
	return &Config{
		Origin: origin,
		Name:   pkgname,
	}, nil
}
