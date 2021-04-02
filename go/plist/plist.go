// Copyright (c) Jeremías Casteglione <jeremias@talkingpts.org>
// See LICENSE file.

// Package plist implements a pkg-plist parser.
package plist

import (
	"bufio"
	"os"
	"path/filepath"
	"strings"

	"uwspkg/manifest"
	"uwspkg/log"
)

type Plist struct {
	m *manifest.Config
	srcdir string
}

func New(m *manifest.Config) *Plist {
	return &Plist{
		m: m,
		srcdir: filepath.FromSlash("/uwspkg/src"),
	}
}

func (p *Plist) Gen(installDir, buildDir string) error {
	fn := filepath.Join(buildDir, "pkg-plist")
	log.Debug("%s gen plist file: %s", p.m.Session, fn)
	fh, err := os.OpenFile(fn, os.O_WRONLY | os.O_CREATE, 0640)
	if err != nil {
		return err
	}
	defer fh.Close()

	// init pkg-plist file
	if err := write(fh, "@owner root"); err != nil {
		return err
	}
	if err := write(fh, "@group root"); err != nil {
		return err
	}
	if err := write(fh, "@mode"); err != nil {
		return err
	}

	// add provided pkg-plist if found
	srcfn := filepath.Join(p.srcdir, p.m.Origin, "pkg-plist")
	log.Debug("%s pkg-plist file: %s", p.m.Session, srcfn)
	if srcfh, err := os.Open(srcfn); err != nil {
		if os.IsNotExist(err) {
			log.Debug("%v", err)
		} else {
			return log.DebugError(err)
		}
	} else {
		defer srcfh.Close()
		x := bufio.NewScanner(srcfh)
		for x.Scan() {
			line := strings.TrimSpace(x.Text())
			xerr := x.Err()
			if xerr != nil {
				return log.DebugError(xerr)
			}
			if err := write(fh, line); err != nil {
				return log.DebugError(err)
			}
		}
	}

	// scan installation dir and add found files (only files, not dirs)
	log.Debug("%s install dir: %s", p.m.Session, installDir)
	plistFiles := func(p string, i os.FileInfo, e error) error {
		if e != nil {
			return log.DebugError(e)
		}
		if !i.IsDir() {
			if err := write(fh, strings.Replace(p, installDir, "", 1)); err != nil {
				return log.DebugError(err)
			}
		}
		return nil
	}
	if err := filepath.Walk(installDir, plistFiles); err != nil {
		return log.DebugError(err)
	}
	return nil
}

func write(fh *os.File, s string) error {
	_, err := fh.WriteString(s+"\n")
	return log.DebugError(err)
}
