// Copyright (c) Jeremías Casteglione <jeremias@talkingpts.org>
// See LICENSE file.

// Package manifest defines the package manifest.
package manifest

import (
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"sync"
	"time"

	"gopkg.in/yaml.v3"

	"uwspkg/libexec"
	"uwspkg/log"
)

type pkgDep struct {
	Origin  string `yaml:"origin" json:"origin"`
	Version string `yaml:"version" json:"version"`
}

type Config struct {
	// internal data
	Package      string    `yaml:"-"`
	Session      string    `yaml:"-"`
	BuildSession string    `yaml:"-"`
	SessionStart time.Time `yaml:"-"`
	OriginPath   string    `yaml:"-"`
	// pkg info
	Origin        string            `yaml:"origin"`
	Name          string            `yaml:"name"`
	Version       string            `yaml:"version"`
	Profile       string            `yaml:"profile"`
	Prefix        string            `yaml:"prefix"`
	Comment       string            `yaml:"comment"`
	Description   string            `yaml:"desc"`
	Licenses      []string          `yaml:"licenses"`
	Maintainer    string            `yaml:"maintainer"`
	WWW           string            `yaml:"www"`
	Categories    []string          `yaml:"categories"`
	ABI           string            `yaml:"abi"`
	PreInstall    string            `yaml:"pre-install"`
	PostInstall   string            `yaml:"post-install"`
	PreDeinstall  string            `yaml:"pre-deinstall"`
	PostDeinstall string            `yaml:"post-deinstall"`
	Users         []string          `yaml:"users"`
	Groups        []string          `yaml:"groups"`
	Deps          map[string]pkgDep `yaml:"deps"`
	// internal pkg info
	Timestamp time.Time `yaml:"-"`
	// actions
	Fetch   string `yaml:"fetch"`
	Depends string `yaml:"depends"`
	Build   string `yaml:"build"`
	Check   string `yaml:"check"`
	Install string `yaml:"install"`
}

func newConfig(origin string) *Config {
	return &Config{Origin: origin, Deps: make(map[string]pkgDep)}
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
	madd := func(k, v interface{}) {
		blob, err := json.Marshal(v)
		if err != nil {
			log.Panic("manifest string encode: %v", err)
		}
		m = fmt.Sprintf("%s%s: %s\n", m, k, blob)
	}
	madd("name", c.Name)
	madd("origin", c.Origin)
	madd("version", c.Version)
	madd("comment", c.Comment)
	madd("maintainer", c.Maintainer)
	madd("www", c.WWW)
	if c.ABI != "" {
		madd("abi", c.ABI)
	}
	madd("prefix", c.Prefix)
	if !c.Timestamp.IsZero() {
		madd("timestamp", c.Timestamp.Unix())
	}
	if len(c.Licenses) == 1 {
		madd("licenselogic", "single")
	}
	madd("licenses", c.Licenses)
	madd("desc", c.Description)
	if len(c.Deps) > 0 {
		madd("deps", c.Deps)
	}
	madd("categories", c.Categories)
	if len(c.Users) > 0 {
		madd("users", c.Users)
	}
	if len(c.Groups) > 0 {
		madd("groups", c.Groups)
	}
	scripts := make(map[string]string)
	if c.PreInstall != "" {
		scripts["pre-install"] = c.PreInstall
	}
	if c.PostInstall != "" {
		scripts["post-install"] = c.PostInstall
	}
	if c.PreDeinstall != "" {
		scripts["pre-deinstall"] = c.PreDeinstall
	}
	if c.PostDeinstall != "" {
		scripts["post-deinstall"] = c.PostDeinstall
	}
	if len(scripts) > 0 {
		madd("scripts", scripts)
	}
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
	var tstamp time.Time
	if st, err := os.Stat(filename); err != nil {
		return log.DebugError(err)
	} else {
		tstamp = st.ModTime()
	}
	blob, err := ioutil.ReadFile(filename)
	if err != nil {
		return log.DebugError(err)
	}
	orig := m.c.Origin
	if err := yaml.Unmarshal(blob, &m.c); err != nil {
		return log.DebugError(err)
	} else {
		if m.c.Origin != orig {
			return fmt.Errorf("%s package origin mismatch: %s", orig, m.c.Origin)
		}
	}
	log.Debug("parse %s", filename)
	m.c.Timestamp = tstamp
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
	c.OriginPath = filepath.FromSlash(c.Origin)
	c.Package = fmt.Sprintf("%s-%s", c.Name, c.Version)
	sess := fmt.Sprintf("%s:%s:%s", time.Now(), orig, c.Profile)
	c.Session = fmt.Sprintf("%x", sha256.Sum256([]byte(sess)))
	c.BuildSession = fmt.Sprintf("uwspkg-build-%s", c.Session)
	if c.Profile == "" {
		c.Profile = "build"
	}
	if c.Prefix == "" {
		c.Prefix = filepath.FromSlash("/uws")
	} else {
		c.Prefix = filepath.FromSlash(c.Prefix)
	}
	if c.Description == "" {
		c.Description = c.Comment
	}
	// actions
	if c.Fetch == "" {
		c.Fetch = "fetch"
	}
	if c.Depends == "" {
		c.Depends = "depends"
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
