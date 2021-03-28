// Copyright (c) Jeremías Casteglione <jeremias@talkingpts.org>
// See LICENSE file.

package manifest

import (
	"fmt"
	"path/filepath"
	"testing"

	_ "uwspkg/_testing/setup"

	. "gopkg.in/check.v1"
)

func Test(t *testing.T) {
	TestingT(t)
}

type TSuite struct {
}

func init() {
	Suite(&TSuite{})
}

func (s *TSuite) TestNewConfig(c *C) {
	m := newConfig("testing")
	c.Check(m.Origin, Equals, "testing")
	c.Check(m.Name, Equals, "")
	c.Check(m.Profile, Equals, "")
	c.Check(m.Session, Equals, "")
	c.Check(len(m.Build), Equals, 0)
}

func (s *TSuite) TestDefaultConfig(c *C) {
	m := New("testdata/load")
	err := m.Load(filepath.FromSlash("testdata/load/manifest.yml"))
	c.Assert(err, IsNil)
	c.Check(m.c.Origin, Equals, "testdata/load")
	c.Check(m.c.Name, Equals, "load")
	c.Check(m.c.Profile, Equals, "build")
	c.Check(len(m.c.Session), Equals, 64)
	c.Check(len(m.c.Source), Equals, 1)
	c.Check(m.c.Source[0], Equals, "./files")
	c.Check(len(m.c.Build), Equals, 1)
	c.Check(m.c.Build[0], Equals, "make")
}

func (s *TSuite) TestBuildScript(c *C) {
	m := New("testdata/build-script")
	err := m.Load(filepath.FromSlash("testdata/build-script/manifest.yml"))
	c.Assert(err, IsNil)
	c.Check(m.c.Origin, Equals, "testdata/build-script")
	c.Check(len(m.c.Build), Equals, 100)
	c.Assert(m.c.Build[0], Equals, "l1")
	c.Assert(m.c.Build[49], Equals, "l50")
	c.Assert(m.c.Build[99], Equals, "l100")
	for i := 0; i < 10; i += 1 {
		m := New("testdata/build-script")
		err := m.Load(filepath.FromSlash("testdata/build-script/manifest.yml"))
		c.Assert(err, IsNil)
		for j := 0; j < 100; j += 1 {
			c.Assert(m.c.Build[j], Equals, fmt.Sprintf("l%d", j+1))
		}
	}
}
