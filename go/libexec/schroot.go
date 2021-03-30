// Copyright (c) Jeremías Casteglione <jeremias@talkingpts.org>
// See LICENSE file.

package libexec

type Chroot struct {
	cmd string
	dir string
	user string
	name string
}

func NewChroot() *Chroot {
	return &Chroot{
		cmd: "/usr/bin/schroot",
		dir: "/root",
		user: "root",
		name: "uwspkg-default",
	}
}

func (c *Chroot) Dirname(d string) {
	c.dir = d
}

func (c *Chroot) User(u string) {
	c.user = u
}

func (c *Chroot) Name(n string) {
	c.name = n
}

func (c *Chroot) Run(env *Env, cmd string, args ...string) error {
	runargs := make([]string, 0)
	runargs = append(runargs, "-p")
	runargs = append(runargs, "-d")
	runargs = append(runargs, c.dir)
	runargs = append(runargs, "-u")
	runargs = append(runargs, c.user)
	runargs = append(runargs, "-c")
	runargs = append(runargs, c.name)
	runargs = append(runargs, "--")
	runargs = append(runargs, cmd)
	for _, x := range args {
		runargs = append(runargs, x)
	}
	return lib.Exec(env, c.cmd, runargs)
}
