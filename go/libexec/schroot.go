// Copyright (c) Jeremías Casteglione <jeremias@talkingpts.org>
// See LICENSE file.

package libexec

import "path"

type Chroot struct {
	name string
	cmd  string
	dir  string
	user string
}

func NewChroot(name string) *Chroot {
	return &Chroot{
		name: name,
		cmd:  "/usr/bin/schroot",
		dir:  "/root",
		user: "",
	}
}

func (c *Chroot) Dirname(d string) {
	c.dir = d
}

func (c *Chroot) User(u string) {
	c.user = u
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
	runargs = append(runargs, "-c")
	runargs = append(runargs, c.name)
	runargs = append(runargs, "--")
	cmdpath := path.Join("/uwspkg/libexec", cmd)
	runargs = append(runargs, cmdpath)
	for _, x := range args {
		runargs = append(runargs, x)
	}
	return lib.Exec(env, c.cmd, runargs)
}
