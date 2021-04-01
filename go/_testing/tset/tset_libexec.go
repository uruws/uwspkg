// Copyright (c) Jeremías Casteglione <jeremias@talkingpts.org>
// See LICENSE file.

package tset

import (
	"fmt"
	"path/filepath"
	"strings"
	"sync"

	"uwspkg/libexec"
)

var _ libexec.Runner = &LibexecMockRunner{}

type LibexecMockRunner struct {
	next      uint
	x         *sync.Mutex
	Calls     map[uint]string
	Commands  map[uint]string
	WithError error
}

func NewLibexecMockRunner() *LibexecMockRunner {
	return &LibexecMockRunner{
		x:        new(sync.Mutex),
		Calls:    make(map[uint]string),
		Commands: make(map[uint]string),
	}
}

func (r *LibexecMockRunner) Exec(env *libexec.Env, cmd string, args []string) error {
	r.x.Lock()
	defer r.x.Unlock()
	r.Calls[r.next] = fmt.Sprintf("%s %s", cmd, strings.Join(args, " "))
	alen := len(args)
	if filepath.Base(cmd) == "schroot" {
		cmd = "schroot"
		aprev := ""
		alen = 0
		acount := false
		for _, a := range args {
			add := false
			if acount {
				alen += 1
			}
			if aprev == "--" {
				cmd += " -- " + a
				acount = true
			} else if aprev == "-n" {
				cmd += " -n "
				add = true
			} else if aprev == "-c" {
				cmd += " -c "
				add = true
			} else if aprev == "-u" && a == "root" {
				cmd += " -u root"
			} else if strings.HasSuffix(aprev, "internal/make") {
				cmd += " " + a
				alen -= 1
			}
			if add {
				if strings.HasPrefix(a, "uwspkg-build-") {
					cmd += "uwspkg-build-ID"
				} else if strings.HasPrefix(a, "build-sess-") {
					cmd += "build-sess-ID"
				} else {
					cmd += a
				}
			}
			if a == "-b" {
				cmd += " -b"
			} else if a == "-e" {
				cmd += " -e"
			}
			aprev = a
		}
	}
	r.Commands[r.next] = fmt.Sprintf("%s [%d]", cmd, alen)
	r.next += 1
	return r.WithError
}

func LibexecRunner(r *LibexecMockRunner) {
	libexec.SetRunner(r)
}

func LibexecDefaultRunner() {
	libexec.SetDefaultRunner()
}
