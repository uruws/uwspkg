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

var defaultConfig *Main = &Main{}

var Files map[int]string = map[int]string{
	0: filepath.FromSlash("/uws/etc/uwspkg.yml"),
	1: filepath.FromSlash("/uws/local/etc/uwspkg.yml"),
	2: filepath.FromSlash("./uwspkg.yml"),
}

const Version uint = 0

type Main struct {
	Libexec              string   `yaml:"-"`
	BuildCfgDir          string   `yaml:"-"`
	BuildEnvPath         string   `yaml:"-"`
	SchrootCfgDir        string   `yaml:"-"`
	Manifest             string   `yaml:"-"`
	Version              uint     `yaml:"version"`
	PkgDir               string   `yaml:"pkgdir"`
	BuildDir             string   `yaml:"build.dir"`
	BuildProfile         []string `yaml:"build.profile"`
	LibexecTimeout       string   `yaml:"libexec.timeout"`
	DebianDeps           []string `yaml:"debian.deps"`
	DebianRepo           string   `yaml:"debian.repo"`
	DebianSecRepo        string   `yaml:"debian.secrepo"`
	DebianInstall        string   `yaml:"debian.install"`
	DebianInstallVariant string   `yaml:"debian.install.variant"`
	DebianDistro         []string `yaml:"debian.distro"`
	PkgBootstrap         string   `yaml:"pkg.bootstrap"`
}

func newMain() *Main {
	bfn := fmt.Sprintf("uwspkg-bootstrap-%s.tgz", )
	return &Main{
		Version:       0,
		PkgDir:        ".",
		Manifest:      "manifest.yml",
		BuildEnvPath:  "/bin:/usr/bin:/usr/sbin",
		SchrootCfgDir: filepath.FromSlash("/etc/schroot"),
		PkgBootstrap:  "20210324",
		Libexec:       defaultConfig.Libexec,
		BuildCfgDir:   defaultConfig.BuildCfgDir,
	}
}

func (c *Main) GetEnviron() map[string]string {
	bfn := fmt.Sprintf("uwspkg-bootstrap-%s.tgz", c.PkgBootstrap)
	return map[string]string{
		"UWSPKG_LIBEXEC": c.Libexec,
		"UWSPKG_CONFIG_VERSION": fmt.Sprintf("%d", c.Version),
		"UWSPKG_CONFIG_SRC": c.BuildCfgDir,
		"UWSPKG_CONFIG_DST": c.SchrootCfgDir,
		"UWSPKG_MANIFEST": c.Manifest,
		"UWSPKG_SOURCE": c.PkgDir,
		"UWSPKG_BUILDDIR": c.BuildDir,
		"UWSPKG_BOOTSTRAP": filepath.Join(c.PkgDir, "build", bfn),
	}
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
	if c.BuildDir == "" {
		c.BuildDir = filepath.FromSlash("/srv/uwspkg")
	} else {
		c.BuildDir, err = filepath.Abs(filepath.Clean(c.BuildDir))
		if err != nil {
			return err
		}
	}
	if len(c.BuildProfile) == 0 {
		c.BuildProfile = []string{"default"}
	}
	if len(c.DebianDeps) == 0 {
		c.DebianDeps = []string{
			"schroot",
			"debootstrap",
			"rsync",
		}
	}
	if c.DebianRepo == "" {
		c.DebianRepo = "http://deb.debian.org/debian"
	}
	if c.DebianSecRepo == "" {
		c.DebianSecRepo = "http://security.debian.org/debian-security"
	}
	if c.DebianInstallVariant == "" {
		c.DebianInstallVariant = "minbase"
	}
	cacheDir := filepath.Join(c.BuildDir, "cache", "debootstrap")
	if c.DebianInstall == "" {
		c.DebianInstall = fmt.Sprintf(
			"debootstrap --force-check-gpg --variant=%s --cache-dir=%s",
			c.DebianInstallVariant, cacheDir)
	}
	if len(c.DebianDistro) == 0 {
		c.DebianDistro = []string{"testing"}
	}
	return nil
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
	loaded := false
	for idx := 0; idx < flen; idx += 1 {
		fn := Files[idx]
		if err := m.LoadFile(fn); err != nil {
			if !os.IsNotExist(err) {
				return nil, err
			}
		} else {
			loaded = true
		}
	}
	if !loaded {
		if err := m.Parse(m.c); err != nil {
			return nil, err
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
