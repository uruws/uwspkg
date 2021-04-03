// Copyright (c) Jeremías Casteglione <jeremias@talkingpts.org>
// See LICENSE file.

// Package plist implements a pkg-plist parser.
package plist

import (
	"bufio"
	"os"
	"path/filepath"
	"strings"

	"uwspkg/log"
	"uwspkg/manifest"
)

type Plist struct {
	m      *manifest.Config
	//~ srcdir string
}

func New(m *manifest.Config) *Plist {
	return &Plist{
		m:      m,
		//~ srcdir: filepath.FromSlash("/uwspkg/src"),
	}
}

func (p *Plist) Gen(installDir, buildDir string) error {
	fn := filepath.Join(buildDir, "pkg-plist")
	log.Debug("%s gen plist file: %s", p.m.Session, fn)
	fh, err := os.OpenFile(fn, os.O_WRONLY|os.O_CREATE, 0640)
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

	done := make(map[string]bool)

	// add manifest plist if not empty
	x := bufio.NewScanner(strings.NewReader(p.m.Plist))
	for x.Scan() {
		line := strings.TrimSpace(x.Text())
		xerr := x.Err()
		if xerr != nil {
			return log.DebugError(xerr)
		}
		if err := write(fh, line); err != nil {
			return log.DebugError(err)
		}
		for _, fn := range p.getFiles(line) {
			if done[fn] {
				return log.NewError("%s: plist duplicate '%s' entry", p.m.Origin, fn)
			}
			done[fn] = true
		}
	}

	// scan installation dir and add found files (only files, not dirs)
	log.Debug("%s install dir: %s", p.m.Session, installDir)
	plistFiles := func(path string, i os.FileInfo, e error) error {
		if e != nil {
			return log.DebugError(e)
		}
		if !i.IsDir() {
			path = strings.Replace(path, installDir, "", 1)
			if strings.HasPrefix(path, p.m.Prefix) {
				path = strings.Replace(path, p.m.Prefix, "", 1)
				path = strings.Replace(path, string(filepath.Separator), "", 1)
			}
			if !done[path] {
				if err := write(fh, path); err != nil {
					return log.DebugError(err)
				}
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
	_, err := fh.WriteString(s + "\n")
	return log.DebugError(err)
}

func (p *Plist) getFiles(line string) []string {
	fl := make([]string, 0)
	i := strings.Split(line, " ")
	if strings.HasPrefix(i[0], "@dir") {
	} else if i[0] == "@owner" {
	} else if i[0] == "@group" {
	} else if i[0] == "@mode" {
	} else if strings.HasPrefix(i[0], "@(") {
		for _, n := range i[1:] {
			fl = append(fl, n)
		}
	} else if !strings.HasPrefix(i[0], "@") {
		for _, n := range i {
			fl = append(fl, n)
		}
	}
	return fl
}
