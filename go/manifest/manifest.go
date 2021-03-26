// Copyright (c) Jeremías Casteglione <jeremias@talkingpts.org>
// See LICENSE file.

// Package manifest defines the package manifest.
package manifest

import (
	"io/ioutil"
	"sync"

	"gopkg.in/yaml.v3"

	"uwspkg/log"
)

type Config struct {
	Origin string
	Name   string
}

func newConfig() *Config {
	return &Config{}
}

type Manifest struct {
	c *Config
	x *sync.Mutex
}

func New() *Manifest {
	return &Manifest{
		c: newConfig(),
		x: new(sync.Mutex),
	}
}

func (m *Manifest) Load(filename string) error {
	log.Debug("load %s", filename)
	m.x.Lock()
	defer m.x.Unlock()
	blob, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Debug("%v", err)
		return err
	}
	if err := yaml.Unmarshal(blob, &m.c); err != nil {
		log.Debug("%v", err)
		return err
	}
	log.Debug("parse %s", filename)
	return m.Parse(m.c)
}

func (m *Manifest) Parse(c *Config) error {
	return nil
}
