// Copyright (c) Jeremías Casteglione <jeremias@talkingpts.org>
// See LICENSE file.

// Package config implements a yaml config manager.
package config

import (
	"io/ioutil"
	"fmt"
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
	Version uint `yaml:version`
}

type Manager struct {
	x *sync.Mutex
	c *Config
}

func New() *Manager {
	return &Manager{
		x: new(sync.Mutex),
		c: new(Config),
	}
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
	return nil
}
