// Copyright (c) Jeremías Casteglione <jeremias@talkingpts.org>
// See LICENSE file.

package build

import (
	"testing"

	_ "uwspkg/_testing/setup"

	"uwspkg/config"
	"uwspkg/_testing/tset"

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
	c.Assert(len(s.mockRunner.Calls), Equals, 2)
	c.Assert(s.mockRunner.Calls, DeepEquals, map[uint]string{
		0: "/uws/libexec/uwspkg/build/setup /srv/uwspkg /uws/etc/schroot /etc/schroot",
		1: "/uws/libexec/uwspkg/build/debian-install /srv/uwspkg minbase http://deb.debian.org/debian http://security.debian.org/debian-security testing",
	})
}

func (s *TSuite) TestEnvSetUpConfig(c *C) {
	s.cfg.BuildProfile = []string{
		"default",
		"build",
	}
	err := EnvSetUp(s.cfg)
	c.Assert(err, IsNil)
	c.Assert(len(s.mockRunner.Calls), Equals, 4)
	c.Assert(s.mockRunner.Calls, DeepEquals, map[uint]string{
		0: "/uws/libexec/uwspkg/build/setup /srv/uwspkg /uws/etc/schroot /etc/schroot",
		1: "/uws/libexec/uwspkg/build/setup-profile /srv/uwspkg /uws/etc/schroot /etc/schroot build",
		2: "/uws/libexec/uwspkg/build/debian-install /srv/uwspkg minbase http://deb.debian.org/debian http://security.debian.org/debian-security testing",
		3: "/uws/libexec/uwspkg/build/debian-install-profile /srv/uwspkg /uws/etc/schroot build",
	})
}
