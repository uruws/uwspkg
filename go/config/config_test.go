// Copyright (c) Jeremías Casteglione <jeremias@talkingpts.org>
// See LICENSE file.

package config

import (
	"testing"

	"uwspkg/log"

	. "gopkg.in/check.v1"
)

func Test(t *testing.T) {
	TestingT(t)
}

type TSuite struct {
}

func init() {
	log.Init("testing")
	Suite(&TSuite{})
}

func (s *TSuite) TestDefaults(c *C) {
	c.Check(Files, HasLen, 3)
}

func (s *TSuite) TestDefaultConfig(c *C) {
	m := newManager()
	c.Check(m.c.Version, Equals, uint(0))
	c.Check(m.c.PkgDir, Equals, ".")
	c.Check(m.c.Manifest, Equals, "manifest.yml")
}
