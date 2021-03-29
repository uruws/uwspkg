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

	"uwspkg/config"
	"uwspkg/libexec"
	"uwspkg/log"
)

type Config struct {
	cfg          *config.Main
	Package      string    `yaml:"-"`
	Session      string    `yaml:"-"`
	BuildSession string    `yaml:"-"`
	SessionStart time.Time `yaml:"-"`
	Origin       string    `yaml:"origin"`
	Name         string    `yaml:"name"`
	Version      string    `yaml:"version"`
	Profile      string    `yaml:"profile"`
	Fetch        string    `yaml:"fetch"`
	Build        string    `yaml:"build"`
	Install      string    `yaml:"install"`
}

func newConfig(cfg *config.Main, origin string) *Config {
	return &Config{cfg: cfg, Origin: origin}
}

func (c *Config) Environ() *libexec.Env {
	e := libexec.NewEnv()
	e.Set("UWSPKG_VERSION_NAME", c.Package)
	e.Set("UWSPKG_BUILD_SESSION", c.BuildSession)
	e.Set("UWSPKG_ORIGIN", c.Origin)
	e.Set("UWSPKG_NAME", c.Name)
	e.Set("UWSPKG_VERSION", c.Version)
	e.Set("UWSPKG_PROFILE", c.Profile)
	return e
}

type Manifest struct {
	c *Config
	x *sync.Mutex
}

func New(cfg *config.Main, origin string) *Manifest {
	return &Manifest{
		c: newConfig(cfg, origin),
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
	if c.Version == "" {
		return fmt.Errorf("%s: empty package version", orig)
	}
	c.Package = fmt.Sprintf("%s-%s", c.Name, c.Version)
	sess := fmt.Sprintf("%s:%s:%s", time.Now(), orig, c.Profile)
	c.Session = fmt.Sprintf("%x", sha256.Sum256([]byte(sess)))
	c.BuildSession = fmt.Sprintf("uwspkg-build-%s", c.Session)
	if c.Profile == "" {
		c.Profile = "build"
	}
	if c.Fetch == "" {
		c.Fetch = "make fetch"
	}
	if c.Build == "" {
		c.Build = "make"
	}
	if c.Install == "" {
		c.Install = "make install"
	}
	return nil
}
