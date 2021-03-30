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
	return nil
}
