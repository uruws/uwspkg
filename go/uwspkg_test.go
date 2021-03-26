// Copyright (c) Jeremías Casteglione <jeremias@talkingpts.org>
// See LICENSE file.

package uwspkg

import (
	"testing"

	_ "uwspkg/_testing/setup"

	"uwspkg/config"

	. "gopkg.in/check.v1"
)

func TestConfigDefaults(t *testing.T) {
	if len(config.ConfigFiles) != 1 {
		t.Fatalf("number of config files: got '%d' - expect '%d'", len(config.ConfigFiles), 1)
	}
}

func Test(t *testing.T) {
	TestingT(t)
}

type TSuite struct {
}

func init() {
	Suite(&TSuite{})
}

func (s *TSuite) TestPackage(c *C) {
	pkg := New("testing/package")
	c.Assert(pkg.orig, Equals, "testing/package")
}

func (s *TSuite) TestPackageLoad(c *C) {
	pkg := New("testing/package_load")
	c.Assert(pkg.orig, Equals, "testing/package_load")
	err := pkg.Load()
	c.Assert(err, IsNil)
}
