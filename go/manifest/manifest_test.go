// Copyright (c) Jeremías Casteglione <jeremias@talkingpts.org>
// See LICENSE file.

package manifest

import (
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
}

func (s *TSuite) TestLoad(c *C) {
	m := New("testdata/load")
	err := m.Load(filepath.FromSlash("testdata/load/manifest.yml"))
	c.Assert(err, IsNil)
	c.Check(m.c.Origin, Equals, "testdata/load")
	c.Check(m.c.Name, Equals, "load")
}
