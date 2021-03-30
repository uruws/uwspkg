// Copyright (c) Jeremías Casteglione <jeremias@talkingpts.org>
// See LICENSE file.

package manifest

import (
	"path/filepath"
	"testing"

	_ "uwspkg/_testing/setup"

	. "gopkg.in/check.v1"
)

func init() {
	Suite(&TSuite{})
}

func Test(t *testing.T) {
	TestingT(t)
}

type TSuite struct {
}

func (s *TSuite) TestNewConfig(c *C) {
	m := newConfig("testing")
	c.Check(m.Origin, Equals, "testing")
	c.Check(m.Name, Equals, "")
	c.Check(m.Version, Equals, "")
	c.Check(m.Profile, Equals, "")
	c.Check(m.Session, Equals, "")
	c.Check(m.Fetch, Equals, "")
	c.Check(m.Build, Equals, "")
	c.Check(m.Check, Equals, "")
	c.Check(m.Install, Equals, "")
}

func (s *TSuite) TestDefaultConfig(c *C) {
	m := New("testdata/load")
	err := m.Load(filepath.FromSlash("testdata/load/manifest.yml"))
	c.Assert(err, IsNil)
	c.Check(m.c.Origin, Equals, "testdata/load")
	c.Check(m.c.Name, Equals, "load")
	c.Check(m.c.Version, Equals, "0")
	c.Check(m.c.Profile, Equals, "build")
	c.Check(len(m.c.Session), Equals, 64)
	c.Check(m.c.Fetch, Equals, "fetch")
	c.Check(m.c.Build, Equals, "build")
	c.Check(m.c.Check, Equals, "check")
	c.Check(m.c.Install, Equals, "install")
}
