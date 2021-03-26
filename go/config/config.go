// Copyright (c) Jeremías Casteglione <jeremias@talkingpts.org>
// See LICENSE file.

// Package config implements a yaml config manager.
package config

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"sync"

	"gopkg.in/yaml.v3"

	"uwspkg/log"
)

var Files map[int]string = map[int]string{
	0: filepath.FromSlash("/uws/etc/uwspkg.yml"),
	1: filepath.FromSlash("/uws/local/etc/uwspkg.yml"),
	2: filepath.FromSlash("./uwspkg.yml"),
}

const Version uint = 0

type Main struct {
	Version       uint   `yaml:"version"`
	PkgDir        string `yaml:"pkgdir"`
	Manifest      string `yaml:"manifest"`
	SchrootCfgDir string `yaml:"schroot.cfgdir"`
}

func (m *manager) Parse(c *Main) error {
	var err error
	c.PkgDir, err = filepath.Abs(filepath.Clean(c.PkgDir))
	if err != nil {
		return err
	}
	if c.Manifest == "" {
		c.Manifest = "manifest.yml"
	}
	if c.SchrootCfgDir == "" {
		c.SchrootCfgDir = filepath.FromSlash("/etc/schroot")
	} else {
		c.SchrootCfgDir, err = filepath.Abs(filepath.Clean(c.SchrootCfgDir))
		if err != nil {
			return err
		}
	}
	return nil
}

func newMain() *Main {
	return &Main{
		Version:  0,
		PkgDir:   ".",
		Manifest: "manifest.yml",
	}
}

type manager struct {
	x *sync.Mutex
	c *Main
}

func newManager() *manager {
	return &manager{
		x: new(sync.Mutex),
		c: newMain(),
	}
}

func Load() (*Main, error) {
	log.Debug("load")
	m := newManager()
	flen := len(Files)
	for idx := 0; idx < flen; idx += 1 {
		fn := Files[idx]
		if err := m.LoadFile(fn); err != nil {
			if !os.IsNotExist(err) {
				return nil, err
			}
		}
	}
	return m.c, nil
}

func (m *manager) LoadFile(name string) error {
	log.Debug("load file: %s", name)
	m.x.Lock()
	defer m.x.Unlock()
	blob, err := ioutil.ReadFile(name)
	if err != nil {
		log.Debug("%v", err)
		return err
	}
	if err := yaml.Unmarshal(blob, &m.c); err != nil {
		return err
	} else {
		if m.c.Version > Version {
			return fmt.Errorf("config invalid version: %d", m.c.Version)
		}
	}
	log.Debug("parse %s", name)
	return m.Parse(m.c)
}
