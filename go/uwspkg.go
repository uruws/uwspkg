// Copyright (c) Jeremías Casteglione <jeremias@talkingpts.org>
// See LICENSE file.

// Package uwspkg defines the package specs.
package uwspkg

import (
	"path"
	"path/filepath"

	"uwspkg/build"
	"uwspkg/config"
	"uwspkg/log"
	"uwspkg/manifest"
)

type Package struct {
	cfg  *config.Main
	orig string
	man  *manifest.Manifest
}

func New(origin string, cfg *config.Main) *Package {
	return &Package{
		cfg:  cfg,
		orig: origin,
		man:  manifest.New(origin),
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
	log.Debug("pkg manifest: %s", pkgman)
	return p.man.Load(pkgman)
}

func (p *Package) Build() error {
	m := p.man.Config()
	if err := build.SetUp(m); err != nil {
		return err
	}
	if err := build.Package(m); err != nil {
		return err
	}
	if err := build.TearDown(m); err != nil {
		return err
	}
	return nil
}
