// Copyright (c) Jeremías Casteglione <jeremias@talkingpts.org>
// See LICENSE file.

package manifest

import (
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
	m := newConfig()
	c.Check(m.Origin, Equals, "")
	c.Check(m.Name, Equals, "")
}
