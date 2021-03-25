// Copyright (c) Jeremías Casteglione <jeremias@talkingpts.org>
// See LICENSE file.

// Package uwspkg defines the package specs.
package uwspkg

type Package struct {
	Name string `yaml:"name"`
}

func New() *Package {
	return &Package{}
}
