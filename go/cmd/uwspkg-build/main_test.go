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
		0: "schroot -u root -c internal-uwspkg -- /uwspkg/libexec/internal/profile-create [2]",
		1: "schroot -c uwspkg-build-ID -- /uwspkg/libexec/internal/make-fetch [1]",
		2: "schroot -c uwspkg-build-ID -- /uwspkg/libexec/internal/source-archive [0]",
		3: "schroot -c uwspkg-build-ID -n build-sess-ID -b [0]",
		4: "schroot -c build-sess-ID -- /uwspkg/libexec/internal/make [1]",
		5: "schroot -c build-sess-ID -- /uwspkg/libexec/internal/make [1]",
		6: "schroot -c build-sess-ID -- /uwspkg/libexec/internal/make-install [1]",
		7: "schroot -c build-sess-ID -e [0]",
		8: "schroot -c internal-uwspkg -- /uwspkg/libexec/internal/make-package [0]",
		9: "schroot -u root -c internal-uwspkg -- /uwspkg/libexec/internal/profile-remove [0]",
	})
}
