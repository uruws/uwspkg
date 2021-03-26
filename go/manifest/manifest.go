// Copyright (c) Jeremías Casteglione <jeremias@talkingpts.org>
// See LICENSE file.

// Package manifest defines the package manifest.
package manifest

import (
	"fmt"
	"io/ioutil"
	"sync"

	"gopkg.in/yaml.v3"

	"uwspkg/log"
)

type Config struct {
	Origin string
	Name   string
}

func newConfig(origin string) *Config {
	return &Config{Origin: origin}
}

type Manifest struct {
	c *Config
	x *sync.Mutex
}

func New(origin string) *Manifest {
	return &Manifest{
		c: newConfig(origin),
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
	orig := m.c.Origin
	if err := yaml.Unmarshal(blob, &m.c); err != nil {
		log.Debug("%v", err)
		return err
	} else {
		if m.c.Origin != orig {
			return fmt.Errorf("%s package origin mismatch: %s", orig, m.c.Origin)
		}
	}
	log.Debug("parse %s", filename)
	return m.Parse(m.c)
}

func (m *Manifest) Parse(c *Config) error {
	return nil
}
