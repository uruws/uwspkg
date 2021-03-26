// Copyright (c) Jeremías Casteglione <jeremias@talkingpts.org>
// See LICENSE file.

// Package uwspkg defines the package specs.
package uwspkg

type Package struct {
	Origin string `yaml:"origin"`
}

func New(origin string) *Package {
	return &Package{Origin: origin}
}
