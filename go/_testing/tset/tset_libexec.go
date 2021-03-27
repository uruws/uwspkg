// Copyright (c) Jeremías Casteglione <jeremias@talkingpts.org>
// See LICENSE file.

package tset

import (
	"uwspkg/libexec"
)

var _ libexec.Runner = &LibexecMockRunner{}

type LibexecMockRunner struct {
}

func newLibexecMockRunner() *LibexecMockRunner {
	return &LibexecMockRunner{}
}

func LibexecSetMockRunner() {
	libexec.SetRunner(newLibexecMockRunner())
}

func LibexecSetDefaultRunner() {
	libexec.SetDefaultRunner()
}
