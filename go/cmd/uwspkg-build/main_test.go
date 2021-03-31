// Copyright (c) Jeremías Casteglione <jeremias@talkingpts.org>
// See LICENSE file.

package main

import (
	"testing"

	"uwspkg"
	"uwspkg/config"
	"uwspkg/_testing/tset"

	. "gopkg.in/check.v1"
)

func Test(t *testing.T) {
	TestingT(t)
}

type TSuite struct {
	cfg *config.Main
	mockRunner *tset.LibexecMockRunner
}

func init() {
	Suite(&TSuite{})
}

func (s *TSuite) SetUpTest(c *C) {
	var err error
	s.cfg, err = config.Load()
	c.Assert(err, IsNil)
	s.mockRunner = tset.NewLibexecMockRunner()
	tset.LibexecRunner(s.mockRunner)
}

func (s *TSuite) TearDownTest(c *C) {
	s.cfg = nil
	s.mockRunner = nil
}

func (s *TSuite) TestBuildPackage(c *C) {
	pkg := uwspkg.New("devel/uwspkg-build", s.cfg)
	err := pkg.Load()
	c.Assert(err, IsNil)
	err = pkg.Build()
	c.Assert(err, IsNil)
	c.Assert(s.mockRunner.Commands, DeepEquals, map[uint]string{
		0: "/usr/bin/schroot [11]",
		1: "/usr/bin/schroot [8]",
		2: "/usr/bin/schroot [7]",
		3: "/usr/bin/schroot [8]",
		4: "/usr/bin/schroot [8]",
		5: "/usr/bin/schroot [8]",
		6: "/usr/bin/schroot [7]",
		7: "/usr/bin/schroot [9]",
	})
}
