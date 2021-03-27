// Copyright (c) Jeremías Casteglione <jeremias@talkingpts.org>
// See LICENSE file.

// Package libexec implements external executable utils.
package libexec

import (
	"fmt"
	"path"
	"path/filepath"
	"time"

	"uwspkg/config"
	"uwspkg/log"
)

var cfg *Config

type Config struct {
	Dir     string
	Timeout time.Duration
}

func init() {
	cfg = &Config{
		Dir: filepath.FromSlash("/uws/libexec/uwspkg"),
		Timeout: 3 * time.Minute,
	}
}

func Configure(c *config.Main) error {
	log.Debug("configure")
	var err error
	if c.Libexec != "" {
		cfg.Dir = c.Libexec
	}
	if c.LibexecTimeout != "" {
		cfg.Timeout, err = time.ParseDuration(c.LibexecTimeout)
		if err != nil {
			return err
		}
	}
	log.Debug("dir: %s", cfg.Dir)
	log.Debug("timeout: %s", cfg.Timeout)
	return nil
}

func Run(cmd string, args ...string) error {
	log.Debug("run: %s %#v", cmd, args)
	cmd = filepath.FromSlash(path.Clean(cmd))
	if filepath.IsAbs(cmd) {
		return fmt.Errorf("cmd should be a relative path: %s", cmd)
	}
	return nil
}
