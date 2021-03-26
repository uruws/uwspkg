// Copyright (c) Jeremías Casteglione <jeremias@talkingpts.org>
// See LICENSE file.

package manifest_test

import (
	"testing"

	_ "uwspkg/_testing/setup"

	"uwspkg/manifest"

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

func (s *TSuite) TestNew(c *C) {
	m := manifest.New()
	c.Check(m.Origin, Equals, "")
	c.Check(m.Name, Equals, "")
}
