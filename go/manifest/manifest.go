// Copyright (c) Jeremías Casteglione <jeremias@talkingpts.org>
// See LICENSE file.

// Package manifest defines the package manifest.
package manifest

import (
	"crypto/sha256"
	"fmt"
	"io/ioutil"
	"sync"
	"time"

	"gopkg.in/yaml.v3"

	"uwspkg/log"
)

type Config struct {
	Origin  string `yaml:"origin"`
	Name    string `yaml:"name"`
	Profile string `yaml:"profile"`
	Session string `yaml:"-"`
}

func newConfig(origin string) *Config {
	return &Config{Origin:  origin}
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

func (m *Manifest) Config() *Config {
	return m.c
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
	orig := c.Origin
	if c.Name == "" {
		return fmt.Errorf("%s: empty package name", orig)
	}
	if c.Profile == "" {
		c.Profile = "build"
	}
	sess := fmt.Sprintf("%s:%s:%s", time.Now(), orig, c.Profile)
	c.Session = fmt.Sprintf("%x", sha256.Sum256([]byte(sess)))
	return nil
}
