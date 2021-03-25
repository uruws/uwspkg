// Copyright (c) Jeremías Casteglione <jeremias@talkingpts.org>
// See LICENSE file.

package logger

import (
	gfmt "fmt"
	"io"
	"log"
	"sync"
)

type Level int

const (
	PANIC Level = iota
	FATAL
	ERROR
	WARN
	MSG
	INFO
	DEBUG
	cReset
)

var levelTag = map[Level]string{
	PANIC: "[PANIC] ",
	FATAL: "[FATAL] ",
	ERROR: "[ERROR] ",
	WARN:  "[WARNING] ",
	MSG:   "",
	INFO:  "",
	DEBUG: "",
}

type Logger struct {
	*sync.Mutex
	log     *log.Logger
	depth   int
	colored bool
	out     io.Writer
	debug   bool
}

func New() *Logger {
	gol := log.New(log.Writer(), "", log.LstdFlags)
	l := &Logger{
		Mutex: new(sync.Mutex),
		log:   gol,
		depth: 2,
	}
	l.out = l.log.Writer()
	log.SetOutput(l)
	return l
}

func (l *Logger) Flags() int {
	return l.log.Flags()
}

func (l *Logger) Prefix() string {
	return l.log.Prefix()
}

func (l *Logger) SetDepth(n int) {
	l.Lock()
	defer l.Unlock()
	l.depth = n + 2
}

func (l *Logger) SetDebug(v bool) {
	l.Lock()
	defer l.Unlock()
	l.debug = v
}

func (l *Logger) SetOutput(out io.Writer) {
	l.Lock()
	defer l.Unlock()
	l.out = nil
	l.log.SetOutput(out)
	l.out = out
}

func (l *Logger) SetFlags(f int) {
	l.Lock()
	defer l.Unlock()
	log.SetFlags(f)
	l.log.SetFlags(f)
}

func (l *Logger) SetPrefix(p string) {
	l.Lock()
	defer l.Unlock()
	log.SetPrefix(p)
	l.log.SetPrefix(p)
}

func (l *Logger) Write(d []byte) (int, error) {
	if l.debug {
		return l.out.Write(d)
	}
	return len(d), nil
}

func (l *Logger) Output(d int, s string) error {
	return l.log.Output(d+2, s)
}

func (l *Logger) tag(lvl Level, msg string) string {
	tag := levelTag[lvl]
	return tag + msg
}

func (l *Logger) Printf(lvl Level, fmt string, args ...interface{}) {
	msg := gfmt.Sprintf(fmt, args...)
	if l.colored {
		l.log.Output(l.depth, l.color(lvl, msg))
	} else {
		l.log.Output(l.depth, l.tag(lvl, msg))
	}
}
