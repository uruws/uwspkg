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
	calls     map[uint]string
	next      uint
	x         *sync.Mutex
	WithError error
}

func NewLibexecMockRunner() *LibexecMockRunner {
	return &LibexecMockRunner{
		calls: make(map[uint]string),
		x:     new(sync.Mutex),
	}
}

func (r *LibexecMockRunner) Exec(cmd string, args []string) error {
	r.x.Lock()
	defer r.x.Unlock()
	r.calls[r.next] = fmt.Sprintf("%s %s", cmd, strings.Join(args, " "))
	r.next += 1
	return r.WithError
}

func LibexecRunner(r *LibexecMockRunner) {
	libexec.SetRunner(r)
}

func LibexecDefaultRunner() {
	libexec.SetDefaultRunner()
}
