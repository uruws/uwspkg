// Copyright (c) Jeremías Casteglione <jeremias@talkingpts.org>
// See LICENSE file.

// Package uwspkg defines the package specs.
package uwspkg

import (
	"path"
	"path/filepath"

	"uwspkg/config"
	"uwspkg/log"
	"uwspkg/manifest"
)

type Package struct {
	cfg  *config.Main
	orig string
	man  *manifest.Config
}

func New(origin string, cfg *config.Main) *Package {
	return &Package{
		cfg:  cfg,
		orig: origin,
	}
}

func (p *Package) Load() error {
	pkgdir, pkgname := path.Split(p.orig)
	pkgdir = filepath.Join(filepath.Clean(filepath.FromSlash(pkgdir)), pkgname)
	log.Debug("cfg pkg dir: %s", p.cfg.PkgDir)
	pkgdir = filepath.Join(p.cfg.PkgDir, pkgdir)
	log.Debug("pkg dir: %s", pkgdir)
	log.Debug("pkg name: %s", pkgname)
	pkgman := filepath.Join(pkgdir, p.cfg.Manifest)
	log.Debug("pkg name: %s", pkgman)
	p.man = manifest.New()
	return nil
}
