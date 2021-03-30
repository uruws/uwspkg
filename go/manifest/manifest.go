// Copyright (c) Jeremías Casteglione <jeremias@talkingpts.org>
// See LICENSE file.

// Package manifest defines the package manifest.
package manifest

import (
	"crypto/sha256"
	"fmt"
	"io/ioutil"
	"path/filepath"
	"sync"
	"time"

	"gopkg.in/yaml.v3"

	"uwspkg/libexec"
	"uwspkg/log"
)

type Config struct {
	// internal data
	Package      string    `yaml:"-"`
	Session      string    `yaml:"-"`
	BuildSession string    `yaml:"-"`
	SessionStart time.Time `yaml:"-"`
	// pkg info
	Origin       string    `yaml:"origin"`
	Name         string    `yaml:"name"`
	Version      string    `yaml:"version"`
	Profile      string    `yaml:"profile"`
	Prefix       string    `yaml:"prefix"`
	Comment      string    `yaml:"comment"`
	Description  string    `yaml:"desc"`
	Licenses     []string  `yaml:"licenses"`
	Maintainer   string    `yaml:"maintainer"`
	WWW          string    `yaml:"www"`
	Categories   []string  `yaml:"categories"`
	// actions
	Fetch        string    `yaml:"fetch"`
	Build        string    `yaml:"build"`
	Check        string    `yaml:"check"`
	Install      string    `yaml:"install"`
}

func newConfig(origin string) *Config {
	return &Config{Origin: origin}
}

func (c *Config) Environ() *libexec.Env {
	e := libexec.NewEnv()
	// from manifest
	e.Set("UWSPKG_VERSION_NAME", c.Package)
	e.Set("UWSPKG_BUILD_SESSION", c.BuildSession)
	e.Set("UWSPKG_ORIGIN", c.Origin)
	e.Set("UWSPKG_NAME", c.Name)
	e.Set("UWSPKG_VERSION", c.Version)
	e.Set("UWSPKG_PROFILE", c.Profile)
	return e
}

func (c *Config) String() string {
	m := ""
	return m
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
	// pkg info
	orig := c.Origin
	if c.Name == "" {
		return fmt.Errorf("%s: empty package name", orig)
	}
	if c.Version == "" {
		return fmt.Errorf("%s: empty package version", orig)
	}
	if c.Comment == "" {
		return fmt.Errorf("%s: empty package comment", orig)
	}
	if len(c.Licenses) == 0 {
		return fmt.Errorf("%s: empty package licenses", orig)
	}
	if c.Maintainer == "" {
		return fmt.Errorf("%s: empty package maintainer", orig)
	}
	if c.WWW == "" {
		return fmt.Errorf("%s: empty package www", orig)
	}
	if len(c.Categories) == 0 {
		return fmt.Errorf("%s: empty package categories", orig)
	}
	c.Package = fmt.Sprintf("%s-%s", c.Name, c.Version)
	sess := fmt.Sprintf("%s:%s:%s", time.Now(), orig, c.Profile)
	c.Session = fmt.Sprintf("%x", sha256.Sum256([]byte(sess)))
	c.BuildSession = fmt.Sprintf("uwspkg-build-%s", c.Session)
	if c.Profile == "" {
		c.Profile = "build"
	}
	if c.Prefix == "" {
		c.Prefix = filepath.FromSlash("/uws")
	}
	if c.Description == "" {
		c.Description = c.Comment
	}
	// actions
	if c.Fetch == "" {
		c.Fetch = "fetch"
	}
	if c.Build == "" {
		c.Build = "build"
	}
	if c.Check == "" {
		c.Check = "check"
	}
	if c.Install == "" {
		c.Install = "install"
	}
	return nil
}
