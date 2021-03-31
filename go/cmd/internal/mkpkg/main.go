// Copyright (c) Jeremías Casteglione <jeremias@talkingpts.org>
// See LICENSE file.

// Package main implements mkpkg internal cmd.
package main

import (
	"fmt"
	"os"
	"path"
	"path/filepath"

	"uwspkg/log"
	"uwspkg/manifest"
)

func main() {
	log.Init("mkpkg")
	log.Print("mkpkg init")
	log.Debug("%v", os.Environ())
	pkgorig := os.Getenv("UWSPKG_ORIGIN")
	if pkgorig == "" {
		log.Error("UWSPKG_ORIGIN not set")
	}
	pkgname := os.Getenv("UWSPKG_NAME")
	if pkgname == "" {
		log.Error("UWSPKG_NAME not set")
	}
	pkgver := os.Getenv("UWSPKG_VERSION")
	if pkgver == "" {
		log.Error("UWSPKG_VERSION not set")
	}
	mfn := path.Join("/uwspkg/src", pkgorig, "manifest.yml")
	x := manifest.New(pkgorig)
	if err := x.Load(mfn); err != nil {
		log.Fatal("%v", err)
	}
	m := x.Config()
	if m.Origin != pkgorig {
		log.Fatal("invalid manifest origin: %s", m.Origin)
	}
	if m.Name != pkgname {
		log.Fatal("invalid manifest name: %s", m.Name)
	}
	if m.Version != pkgver {
		log.Fatal("invalid manifest version: %s", m.Version)
	}
	log.Print("mkpkg end")
}

func writeManifest(m *manifest.Config) error {
	fn := filepath.Join("/build", m.BuildSession, m.Package+".manifest")
	log.Debug("%s write manifest: %s", m.Session, fn)
	if _, err := os.Stat(fn); err == nil {
		return fmt.Errorf("%s: file already exists", fn)
	} else {
		if !os.IsNotExist(err) {
			return err
		}
	}
	fh, err := os.OpenFile(fn, os.O_WRONLY, 0640)
	if err != nil {
		return err
	}
	defer fh.Close()
	if n, err := fh.WriteString(m.String()); err != nil {
		return err
	} else {
		if n != len(m.String()) {
			return fmt.Errorf("write manifest: wrong number of bytes")
		}
	}
	return nil
}
