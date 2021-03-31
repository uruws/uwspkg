// Copyright (c) Jeremías Casteglione <jeremias@talkingpts.org>
// See LICENSE file.

// Package main implements mkpkg internal cmd.
package main

import (
	"io/ioutil"
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
	log.Print("mkpkg end")
}

func writeManifest(m *manifest.Config, buildDir string) error {
	fn := filepath.Join(buildDir, "+MANIFEST")
	log.Debug("%s write manifest: %s", m.Session, fn)
	return ioutil.WriteFile(fn, []byte(m.String()), 0640)
}
