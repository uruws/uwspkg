// Copyright (c) Jeremías Casteglione <jeremias@talkingpts.org>
// See LICENSE file.

package tset

import (
	"fmt"
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
	r.Commands[r.next] = fmt.Sprintf("%s [%d]", cmd, len(args))
	r.next += 1
	return r.WithError
}

func LibexecRunner(r *LibexecMockRunner) {
	libexec.SetRunner(r)
}

func LibexecDefaultRunner() {
	libexec.SetDefaultRunner()
}
