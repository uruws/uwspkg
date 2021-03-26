// Copyright (c) Jeremías Casteglione <jeremias@talkingpts.org>
// See LICENSE file.

package uwspkg

import (
	"testing"

	_ "uwspkg/_testing/setup"

	"uwspkg/config"

	. "gopkg.in/check.v1"
)

func Test(t *testing.T) {
	TestingT(t)
}

type TSuite struct {
	cfg *config.Main
}

func init() {
	Suite(&TSuite{})
}

func (s *TSuite) SetUpTest(c *C) {
	var err error
	s.cfg, err = config.Load()
	c.Assert(err, IsNil)
}

func (s *TSuite) TearDownTest(c *C) {
	s.cfg = nil
}

func (s *TSuite) TestConfigDefaults(c *C) {
	c.Check(config.Files, HasLen, 2)
}

func (s *TSuite) TestPackage(c *C) {
	pkg := New("testing/package", s.cfg)
	c.Assert(pkg.cfg, Equals, s.cfg)
	c.Assert(pkg.orig, Equals, "testing/package")
}

func (s *TSuite) TestPackageLoad(c *C) {
	pkg := New("testdata/base", s.cfg)
	c.Assert(pkg.orig, Equals, "testdata/base")
	err := pkg.Load()
	c.Assert(err, IsNil)
}
