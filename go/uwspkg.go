// Copyright (c) Jeremías Casteglione <jeremias@talkingpts.org>
// See LICENSE file.

// Package uwspkg defines the package specs.
package uwspkg

import (
	"fmt"
	"os"
	"path"
	"path/filepath"
	"time"

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
	log.Print("Build %s-%s.", p.orig, m.Version)
	log.Debug("build profile: %s", m.Profile)
	// check profile
	profile := filepath.Join(p.cfg.SchrootCfgDir, "uwspkg-"+m.Profile)
	if st, err := os.Stat(profile); err != nil {
		if os.IsNotExist(err) {
			return fmt.Errorf("%s invalid build profile: %s", p.orig, m.Profile)
		} else {
			return err
		}
	} else {
		if !st.IsDir() {
			return fmt.Errorf("%s invalid build profile: %s", p.orig, m.Profile)
		}
	}
	// defer tear down
	defer func() {
		if errlist := build.TearDown(p.cfg, m); len(errlist) > 0 {
			for _, err := range errlist {
				log.Error("%v", err)
			}
			log.Fatal("build tear down failed with %d error(s)", len(errlist))
		}
	}()
	// defer fail report
	failed := true
	defer func() {
		if failed {
			log.Error("Build %s failed.", p.orig)
		}
	}()
	// setup build env
	if err := build.SetUp(p.cfg, m); err != nil {
		return err
	}
	// build source package
	log.Print("Make %s source %s.", m.Origin, m.Source)
	if m.Source == m.Origin {
		if err := build.Source(m); err != nil {
			return err
		}
	} else {
		src := New(m.Source, p.cfg)
		if err := src.Load(); err != nil {
			return err
		}
		sm := src.man.Config()
		sm.Session = m.Session
		if err := build.Source(sm); err != nil {
			return err
		}
	}
	// build package
	if err := build.Package(m); err != nil {
		return err
	}
	// success report
	failed = false
	log.Info("Build %s-%s done in %s.", p.orig, m.Version, time.Now().Sub(m.SessionStart))
	return nil
}
