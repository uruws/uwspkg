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

func (s *TSuite) TestNew(c *C) {
	m, err := New("testing/manifest", "manifest",
		filepath.FromSlash("testing/manifest/manifest.yml"))
	c.Assert(err, IsNil)
	c.Check(m.Origin, Equals, "testing/manifest")
	c.Check(m.Name, Equals, "manifest")
}
