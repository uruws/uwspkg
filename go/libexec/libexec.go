// Copyright (c) Jeremías Casteglione <jeremias@talkingpts.org>
// See LICENSE file.

// Package libexec implements external executable utils.
package libexec

import (
	"context"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"time"

	"uwspkg/config"
	"uwspkg/log"
)

var (
	cfg *Config
	lib Runner
)

type Config struct {
	Dir     string
	Timeout time.Duration
}

var _ Runner = &impl{}

type Runner interface {
	Exec(string, []string) error
}

type impl struct {
}

func (r *impl) Exec(cmdpath string, args []string) error {
	ctx, cancel := context.WithTimeout(context.Background(), cfg.Timeout)
	defer cancel()
	cmd := exec.CommandContext(ctx, cmdpath, args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	log.Debug("exec: %s", cmd.String())
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("%s: %v", cmdpath, err)
	}
	return nil
}

func init() {
	cfg = &Config{
		Dir:     filepath.FromSlash("/uws/libexec/uwspkg"),
		Timeout: 3 * time.Minute,
	}
	SetDefaultRunner()
}

func SetRunner(r Runner) {
	lib = nil
	lib = r
}

func SetDefaultRunner() {
	SetRunner(&impl{})
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

func Run(cmdname string, args ...string) error {
	cmdname = filepath.FromSlash(cmdname)
	log.Debug("run: %s %v", cmdname, args)
	if filepath.IsAbs(cmdname) {
		return fmt.Errorf("cmd should be a relative path: %s", cmdname)
	}
	log.Print("Run %s.", cmdname)
	cmdpath := filepath.Join(cfg.Dir, cmdname)
	log.Debug("cmd path: %s", cmdpath)
	if !strings.HasPrefix(cmdpath, cfg.Dir) {
		return fmt.Errorf("%s: cmd path outside of libexec dir", cmdpath)
	}
	return lib.Exec(cmdpath, args)
}
