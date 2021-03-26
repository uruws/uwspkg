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

var ConfigFiles map[int]string = map[int]string{
	0: filepath.FromSlash("/uws/etc/uwspkg.yml"),
	1: filepath.FromSlash("/uws/local/etc/uwspkg.yml"),
	2: filepath.FromSlash("./uwspkg.yml"),
}

var cfg *Manager

func init() {
	cfg = New()
}

func Load() error {
	log.Debug("load")
	flen := len(ConfigFiles)
	for idx := 0; idx < flen; idx += 1 {
		fn := ConfigFiles[idx]
		if err := cfg.LoadFile(fn); err != nil {
			if !os.IsNotExist(err) {
				return err
			}
		}
	}
	return nil
}

const Version uint = 1

type Config struct {
	Version  uint   `yaml:version`
	PkgDir   string `yaml:"pkgdir"`
	Manifest string `yaml:"manifest"`
}

func newConfig() *Config {
	return &Config{
		Version: 0,
		PkgDir: ".",
		Manifest: "manifest.yml",
	}
}

type Manager struct {
	x *sync.Mutex
	c *Config
}

func New() *Manager {
	return &Manager{
		x: new(sync.Mutex),
		c: newConfig(),
	}
}

func Get() *Config {
	return cfg.c
}

func (m *Manager) LoadFile(name string) error {
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

func (m *Manager) Parse(c *Config) error {
	var err error
	c.PkgDir, err = filepath.Abs(filepath.Clean(c.PkgDir))
	if err != nil {
		return err
	}
	c.Manifest = filepath.Clean(c.Manifest)
	return nil
}
