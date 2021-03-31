// Copyright (c) Jeremías Casteglione <jeremias@talkingpts.org>
// See LICENSE file.

package build

import (
	"testing"

	_ "uwspkg/_testing/setup"

	"uwspkg/_testing/tset"
	"uwspkg/config"

	. "gopkg.in/check.v1"
)

func init() {
	Suite(&TSuite{})
}

func Test(t *testing.T) {
	TestingT(t)
}

type TSuite struct {
	cfg        *config.Main
	mockRunner *tset.LibexecMockRunner
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
	tset.LibexecDefaultRunner()
}

func (s *TSuite) TestEnvSetUp(c *C) {
	err := EnvSetUp(s.cfg)
	c.Assert(err, IsNil)
	c.Assert(len(s.mockRunner.Calls), Equals, 3)
	c.Assert(s.mockRunner.Commands, DeepEquals, map[uint]string{
		0: "/uws/libexec/uwspkg/build/setup [7]",
		1: "/uws/libexec/uwspkg/build/debian-install [6]",
		2: "/uws/libexec/uwspkg/build/debian-install-profile [1]",
	})
}

func (s *TSuite) TestEnvSetUpConfig(c *C) {
	s.cfg.BuildProfile = []string{
		"default",
		"build",
	}
	err := EnvSetUp(s.cfg)
	c.Assert(err, IsNil)
	c.Assert(len(s.mockRunner.Calls), Equals, 5)
	c.Assert(s.mockRunner.Commands, DeepEquals, map[uint]string{
		0: "/uws/libexec/uwspkg/build/setup [7]",
		1: "/uws/libexec/uwspkg/build/setup-profile [4]",
		2: "/uws/libexec/uwspkg/build/debian-install [6]",
		3: "/uws/libexec/uwspkg/build/debian-install-profile [1]",
		4: "/uws/libexec/uwspkg/build/debian-install-profile [1]",
	})
}
