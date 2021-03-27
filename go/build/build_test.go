// Copyright (c) Jeremías Casteglione <jeremias@talkingpts.org>
// See LICENSE file.

package build

import (
	"testing"

	_ "uwspkg/_testing/setup"

	"uwspkg/config"
	"uwspkg/_testing/tset"

	. "gopkg.in/check.v1"
)

func init() {
	Suite(&TSuite{})
}

func Test(t *testing.T) {
	TestingT(t)
}

type TSuite struct {
	cfg        *config.Main
	mockRunner *tset.LibexecMockRunner
}

func (s *TSuite) SetUpTest(c *C) {
	var err error
	s.cfg, err = config.Load()
	c.Assert(err, IsNil)
	s.mockRunner = tset.NewLibexecMockRunner()
	tset.LibexecRunner(s.mockRunner)
}

func (s *TSuite) TearDownTest(c *C) {
	s.cfg = nil
	s.mockRunner = nil
	tset.LibexecDefaultRunner()
}

func (s *TSuite) TestEnvSetUp(c *C) {
	err := EnvSetUp(s.cfg)
	c.Assert(err, IsNil)
}
