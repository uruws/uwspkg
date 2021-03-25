// Copyright (c) Jeremías Casteglione <jeremias@talkingpts.org>
// See LICENSE file.

// Package log provides the logger functionalities.
package log

import (
	"errors"
	"fmt"
	"io"
	"log"
	"os"
	"strings"

	"uwspkg/log/internal/logger"
)

var (
	cdepth     int  = 1
	debug      bool = false
	info       bool = true
	verbose    bool = true
	warn       bool = true
	debugFlags int  = log.Llongfile
	stdFlags   int  = log.Ldate | log.Ltime | log.Lmicroseconds | log.Lmsgprefix
)

var l *logger.Logger

func init() {
	l = logger.New()
	l.SetDepth(cdepth)
	l.SetFlags(stdFlags)
	setVerbose()
}

func setDebugFlags(s string) {
	l.Lock()
	defer l.Unlock()
	switch s {
	case "std":
		debugFlags = stdFlags
		return
	case "all":
		debugFlags = stdFlags | log.Llongfile
		return
	}
	var flags int
	for _, f := range strings.Fields(s) {
		switch f {
		case "date":
			flags = flags | log.Ldate
		case "time":
			flags = flags | log.Ltime
		case "microseconds":
			flags = flags | log.Lmicroseconds
		case "longfile":
			flags = flags | log.Llongfile
		case "shortfile":
			flags = flags | log.Lshortfile
		case "UTC":
			flags = flags | log.LUTC
		}
	}
	if flags == 0 {
		flags = log.Llongfile
	}
	debugFlags = flags
}

func setQuiet() {
	l.SetDebug(false)
	l.SetFlags(stdFlags)
	l.Lock()
	defer l.Unlock()
	debug = false
	info = false
	verbose = false
	warn = false
}

func setDebug() {
	l.SetFlags(debugFlags)
	l.SetDebug(true)
	l.Lock()
	defer l.Unlock()
	debug = true
	info = true
	verbose = true
	warn = true
}

func setInfo() {
	l.SetDebug(false)
	l.SetFlags(stdFlags)
	l.Lock()
	defer l.Unlock()
	debug = false
	info = true
	verbose = false
	warn = true
}

func setVerbose() {
	l.SetDebug(false)
	l.SetFlags(stdFlags)
	l.Lock()
	defer l.Unlock()
	debug = false
	info = true
	verbose = true
	warn = true
}

func setWarning() {
	l.SetDebug(false)
	l.SetFlags(stdFlags)
	l.Lock()
	defer l.Unlock()
	debug = false
	info = false
	verbose = false
	warn = true
}

func setMode(lvl string) {
	switch lvl {
	case "quiet":
		setQuiet()
	case "debug":
		setDebug()
	case "info":
		setInfo()
	case "warn":
		setWarning()
	default:
		setVerbose()
	}
}

func setColors(cfg string) {
	l.SetColors(cfg)
}

func setPrefix(name string) {
	p := fmt.Sprintf("[%s:%d] ", name, os.Getpid())
	l.SetPrefix(p)
}

func setOutput(out io.Writer) {
	l.SetOutput(out)
}

func Init(progname string) {
	setPrefix(progname)
	if mode := os.Getenv("UWSPKG_LOG"); mode != "" {
		setMode(mode)
	} else {
		setMode("default")
	}
	if colors := os.Getenv("UWSPKG_LOG_COLORS"); colors != "" {
		setColors(colors)
	} else {
		setColors("auto")
	}
}

func Output(calldepth int, s string) error {
	return l.Output(calldepth, s)
}

func Panic(format string, v ...interface{}) {
	err := errors.New(fmt.Sprintf(format, v...))
	l.Printf(logger.PANIC, format, v...)
	panic(err)
}

func Print(format string, v ...interface{}) {
	if verbose {
		l.Printf(logger.MSG, format, v...)
	}
}

func Debug(format string, v ...interface{}) {
	if debug {
		l.Printf(logger.DEBUG, format, v...)
	}
}

func Error(format string, v ...interface{}) error {
	err := errors.New(fmt.Sprintf(format, v...))
	l.Printf(logger.ERROR, format, v...)
	return err
}

var osExit func(int) = os.Exit

func Fatal(format string, v ...interface{}) {
	l.Printf(logger.FATAL, format, v...)
	osExit(2)
}

func Warn(format string, v ...interface{}) {
	if warn {
		l.Printf(logger.WARN, format, v...)
	}
}

func Info(format string, v ...interface{}) {
	if info {
		l.Printf(logger.INFO, format, v...)
	}
}
