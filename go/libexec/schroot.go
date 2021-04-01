// Copyright (c) Jeremías Casteglione <jeremias@talkingpts.org>
// See LICENSE file.

package libexec

import (
	"path"

	"uwspkg/log"
)

type Chroot struct {
	name string
	cmd  string
	dir  string
	user string
	sess string
}

func NewChroot(name string) *Chroot {
	return &Chroot{
		name: name,
		cmd:  "/usr/bin/schroot",
		dir:  "/root",
		user: "",
	}
}

func (c *Chroot) Dir(d string) {
	c.dir = d
}

func (c *Chroot) User(u string) {
	c.user = u
}

func (c *Chroot) SessionBegin(sess string) error {
	env := NewEnv()
	args := []string{
		0: "-d",
		1: c.dir,
		2: "-u",
		3: c.user,
		4: "-c",
		5: c.name,
		6: "-n",
		7: sess,
		8: "-b",
	}
	if err := lib.Exec(env, c.cmd, args); err != nil {
		return err
	}
	c.sess = sess
	return nil
}

func (c *Chroot) SessionEnd() {
	if c.sess == "" {
		return
	}
	env := NewEnv()
	args := []string{
		0: "-d",
		1: c.dir,
		2: "-u",
		3: c.user,
		4: "-c",
		5: c.sess,
		6: "-e",
	}
	if err := lib.Exec(env, c.cmd, args); err != nil {
		log.Error("%v", err)
	}
	c.sess = ""
}

func (c *Chroot) Run(env *Env, cmd string, args ...string) error {
	runargs := make([]string, 0)
	runargs = append(runargs, "-p")
	runargs = append(runargs, "-d")
	runargs = append(runargs, c.dir)
	if c.user != "" {
		runargs = append(runargs, "-u")
		runargs = append(runargs, c.user)
	}
	if c.sess != "" {
		runargs = append(runargs, "-r")
		runargs = append(runargs, "-c")
		runargs = append(runargs, c.sess)
	} else {
		runargs = append(runargs, "-c")
		runargs = append(runargs, c.name)
	}
	runargs = append(runargs, "--")
	cmdpath := path.Join("/uwspkg/libexec", cmd)
	runargs = append(runargs, cmdpath)
	for _, x := range args {
		runargs = append(runargs, x)
	}
	return lib.Exec(env, c.cmd, runargs)
}
