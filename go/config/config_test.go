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
	c.Check(ConfigFiles, HasLen, 3)
}

func (s *TSuite) TestDefaultConfig(c *C) {
	cfg := newConfig()
	c.Check(cfg.Version, Equals, uint(0))
	c.Check(cfg.PkgDir, Equals, ".")
}
