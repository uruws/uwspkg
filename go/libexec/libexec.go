// Copyright (c) Jeremías Casteglione <jeremias@talkingpts.org>
// See LICENSE file.

// Package libexec implements external executable utils.
package libexec

import (
	"context"
	"fmt"
	"os"
	"os/exec"
	"os/user"
	"path/filepath"
	"strconv"
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
	Dir           string
	Timeout       time.Duration
	DeveloperMode bool
}

type Env struct {
	d map[string]string
	l []string
}

func NewEnv() *Env {
	e := &Env{
		d: make(map[string]string),
		l: make([]string, 0),
	}
	u, err := user.Current()
	if err != nil {
		log.Panic("%v", err)
	}
	envadd := func(f string, args ...interface{}) {
		e.l = append(e.l, fmt.Sprintf(f, args...))
	}
	envadd("%s=%s", "USER", u.Username)
	envadd("%s=%s", "LOGNAME", u.Username)
	envadd("%s=%s", "HOME", u.HomeDir)
	if term := os.Getenv("TERM"); term != "" {
		envadd("%s=%s", "TERM", term)
	}
	envadd("SHELL=/bin/sh")
	envadd("PATH=/bin:/usr/bin:/usr/local/bin")
	if loglvl := os.Getenv("UWSPKG_LOG"); loglvl == "" {
		envadd("UWSPKG_LOG=default")
	} else {
		envadd("UWSPKG_LOG=%s", loglvl)
	}
	envadd("UWSPKG_DEVELOPER_MODE=%s", strconv.FormatBool(cfg.DeveloperMode))
	return e
}

func (e *Env) getEnviron() []string {
	x := make([]string, 0)
	for k, v := range e.d {
		x = append(x, fmt.Sprintf("%s=%s", k, v))
	}
	// force user settings
	for _, v := range e.l {
		x = append(x, v)
	}
	return x
}

func (e *Env) Set(key, val string) {
	if key == "USER" ||
		key == "LOGNAME" ||
		key == "HOME" ||
		key == "TERM" ||
		key == "SHELL" ||
		key == "PATH" {
		return
	}
	e.d[key] = val
}

var _ Runner = &impl{}

type Runner interface {
	Exec(*Env, string, []string) error
}

type impl struct {
}

func (r *impl) Exec(env *Env, cmdpath string, args []string) error {
	ctx, cancel := context.WithTimeout(context.Background(), cfg.Timeout)
	defer cancel()
	cmd := exec.CommandContext(ctx, cmdpath, args...)
	cmd.Env = env.getEnviron()
	//~ log.Debug("ENV: %v", cmd.Env)
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
	log.Debug("dir: %s", cfg.Dir)
	if c.LibexecTimeout != "" {
		cfg.Timeout, err = time.ParseDuration(c.LibexecTimeout)
		if err != nil {
			return err
		}
	}
	log.Debug("timeout: %s", cfg.Timeout)
	cfg.DeveloperMode = c.DeveloperMode
	log.Debug("developer_mode: %v", cfg.DeveloperMode)
	return nil
}

func EnvConfig(c *config.Main) *Env {
	e := NewEnv()
	for k, v := range c.GetEnviron() {
		e.Set(k, v)
	}
	return e
}

func Run(cmdname string, args ...string) error {
	return RunEnv(NewEnv(), cmdname, args...)
}

func RunEnv(env *Env, cmdname string, args ...string) error {
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
	return lib.Exec(env, cmdpath, args)
}
