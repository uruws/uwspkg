// Copyright (c) Jeremías Casteglione <jeremias@talkingpts.org>
// See LICENSE file.

// Package main implements mkpkg internal cmd.
package main

import (
	"io/ioutil"
	"os"
	"os/exec"
	"path"
	"path/filepath"

	"uwspkg/log"
	"uwspkg/manifest"
)

func main() {
	log.Init("mkpkg")
	log.Print("mkpkg init")
	log.Debug("%v", os.Environ())
	// load env
	buildSess := os.Getenv("UWSPKG_BUILD_SESSION")
	if buildSess == "" {
		log.Fatal("UWSPKG_BUILD_SESSION not set")
	}
	buildDir := os.Getenv("UWSPKG_BUILDDIR")
	if buildDir == "" {
		log.Fatal("UWSPKG_BUILDDIR not set")
	}
	destDir := os.Getenv("UWSPKG_DESTDIR")
	if destDir == "" {
		log.Fatal("UWSPKG_DESTDIR not set")
	}
	pkgorig := os.Getenv("UWSPKG_ORIGIN")
	if pkgorig == "" {
		log.Fatal("UWSPKG_ORIGIN not set")
	}
	pkgname := os.Getenv("UWSPKG_NAME")
	if pkgname == "" {
		log.Fatal("UWSPKG_NAME not set")
	}
	pkgver := os.Getenv("UWSPKG_VERSION")
	if pkgver == "" {
		log.Fatal("UWSPKG_VERSION not set")
	}
	// load manifest
	mfn := path.Join("/uwspkg/src", pkgorig, "manifest.yml")
	x := manifest.New(pkgorig)
	if err := x.Load(mfn); err != nil {
		log.Fatal("%v", err)
	}
	m := x.Config()
	// build session settings
	m.BuildSession = buildSess
	// check laoded manifest
	if m.Origin != pkgorig {
		log.Fatal("invalid manifest origin: %s", m.Origin)
	}
	if m.Name != pkgname {
		log.Fatal("invalid manifest name: %s", m.Name)
	}
	if m.Version != pkgver {
		log.Fatal("invalid manifest version: %s", m.Version)
	}
	// write package +MANIFEST
	if err := writeManifest(m, buildDir); err != nil {
		log.Fatal("%v", err)
	}
	// save generated files
	saveSource(m)
	log.Print("mkpkg end")
}

func writeManifest(m *manifest.Config, buildDir string) error {
	fn := filepath.Join(buildDir, "+MANIFEST")
	log.Debug("%s write manifest: %s", m.Session, fn)
	return ioutil.WriteFile(fn, []byte(m.String()), 0640)
}

func mv(src, dst string) {
	log.Debug("mv %s %s", src, dst)
	cmd := exec.Command("/usr/bin/mv", "-v", src, dst)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		log.Fatal("mv %s %s: %v", src, dst, err)
	}
}

func saveSource(m *manifest.Config) {
	srcfn := filepath.Join("/build", m.BuildSession, m.Package+"-source.tgz")
	dstfn := path.Join("/uwspkg/repo/src", path.Dir(m.Origin),
		m.Package+"-source.tgz")
	log.Debug("save source: %s -> %s", srcfn, dstfn)
	dstd := path.Dir(dstfn)
	if err := os.MkdirAll(dstd, 0750); err != nil {
		log.Fatal("%v", err)
	}
	mv(srcfn, dstfn)
}
