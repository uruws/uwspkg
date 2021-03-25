// Copyright (c) Jeremías Casteglione <jeremias@talkingpts.org>
// See LICENSE file.

// Package config implements a yaml config manager.
package config

import (
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

var cfg *Config
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

type Loader struct {
	Version int `yaml:"version"`
}

type Config struct {
	m *sync.Mutex
	Loader *Loader `yaml:"config"`
}

func New() *Config {
	return &Config{
		m: new(sync.Mutex),
		Loader: &Loader{},
	}
}

func (c *Config) LoadFile(name string) error {
	log.Debug("load file: %s", name)
	c.m.Lock()
	defer c.m.Unlock()
	blob, err := ioutil.ReadFile(name)
	if err != nil {
		log.Debug("%v", err)
		return err
	}
	if err := yaml.Unmarshal(blob, &c); err != nil {
		return err
	} else {
	}
	return nil
}
