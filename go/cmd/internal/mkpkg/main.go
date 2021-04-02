// Copyright (c) Jeremías Casteglione <jeremias@talkingpts.org>
// See LICENSE file.

// Package main implements mkpkg internal cmd.
package main

import (
	"flag"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"

	"uwspkg/log"
	"uwspkg/manifest"
	"uwspkg/plist"
)

// runs from /uwspkg/libexec/internal/mkpkg inside internal-uwspkg chroot

func main() {
	log.Init("mkpkg")
	log.Print("mkpkg init")
	var (
		genManifest string
		genPlist    string
		pkgOrig     string
		buildDir    string
		destDir     string
	)
	flag.StringVar(&genManifest, "manifest", "", "gen manifest")
	flag.StringVar(&genPlist, "plist", "", "gen pkg-plist")
	flag.StringVar(&pkgOrig, "pkg", "", "pkg origin")
	flag.StringVar(&buildDir, "builddir", "", "dir where to save generated files")
	flag.StringVar(&destDir, "destdir", "", "dir from where to generate plist files")
	flag.Parse()
	if genManifest != "" {
		x := manifest.New(pkgOrig)
		if err := x.Load(genManifest); err != nil {
			log.Fatal("%v", err)
		}
		m := x.Config()
		if err := writeManifest(m, buildDir); err != nil {
			log.Fatal("%v", err)
		}
	} else if genPlist != "" {
		x := manifest.New(pkgOrig)
		if err := x.Load(genPlist); err != nil {
			log.Fatal("%v", err)
		}
		m := x.Config()
		p := plist.New(m)
		if err := p.Gen(destDir, buildDir); err != nil {
			log.Fatal("%v", err)
		}
	} else {
		doMain()
	}
	log.Print("mkpkg end")
}

func doMain() {
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
	p := plist.New(m)
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
	// write package metadata
	if err := writeManifest(m, buildDir); err != nil {
		log.Fatal("%v", err)
	}
	if err := p.Gen(destDir, buildDir); err != nil {
		log.Fatal("%v", err)
	}
}

func writeManifest(m *manifest.Config, buildDir string) error {
	fn := filepath.Join(buildDir, "+MANIFEST")
	log.Debug("%s write manifest: %s", m.Session, fn)
	return ioutil.WriteFile(fn, []byte(m.String()), 0640)
}
