// Copyright (c) Jeremías Casteglione <jeremias@talkingpts.org>
// See LICENSE file.

package logger

import (
	"fmt"
	"os"
	"strings"
)

// colors and escape stolen from: golang.org/x/crypto/ssh/terminal/terminal.go
const escape = 27

var (
	reset   = "0"
	black   = "30"
	red     = "31"
	green   = "32"
	yellow  = "33"
	blue    = "34"
	magenta = "35"
	cyan    = "36"
	white   = "37"
	grey    = "1;30"
)

var levelColor = map[Level]string{
	PANIC:  tilt(red),
	FATAL:  bold(red),
	ERROR:  red,
	WARN:   italic(yellow),
	MSG:    magenta,
	INFO:   cyan,
	DEBUG:  green,
	cReset: reset,
}

func bold(v string) string {
	return "01;" + v
}

func italic(v string) string {
	return "03;" + v
}

func underline(v string) string {
	return "04;" + v
}

func tilt(v string) string {
	return "05;" + v
}

func invert(v string) string {
	return "07;" + v
}

func (l *Logger) Colors() bool {
	return l.colored
}

var esc = []byte{escape}

func (l *Logger) color(lvl Level, msg string) string {
	col := levelColor[lvl]
	rst := levelColor[cReset]
	tag := levelTag[lvl]
	return fmt.Sprintf("%s[%sm%s%s%s[%sm", esc, col, tag, msg, esc, rst)
}

func (l *Logger) SetColors(cfg string) {
	l.Lock()
	defer l.Unlock()
	cfg = strings.TrimSpace(cfg)
	l.colored = false
	switch cfg {
	case "":
		return
	case "off":
		return
	case "on":
		l.colored = true
	case "auto":
		if istty(os.Stdout) && istty(os.Stderr) {
			l.colored = true
		}
	default:
		l.colored = true
	}
	if l.colored {
		setColors(cfg)
	}
}

func istty(fh *os.File) bool {
	if st, err := fh.Stat(); err == nil {
		m := st.Mode()
		if m&os.ModeDevice != 0 && m&os.ModeCharDevice != 0 {
			return true
		}
	}
	return false
}

func setColors(cfg string) {
	for _, opt := range strings.Split(cfg, ":") {
		i := strings.Split(opt, "=")
		if len(i) == 2 {
			key := i[0]
			val := i[1]
			if val == "" {
				continue
			}
			// TODO: validate value?
			switch key {
			case "rst":
				levelColor[cReset] = val
			case "pnc":
				levelColor[PANIC] = val
			case "ftl":
				levelColor[FATAL] = val
			case "err":
				levelColor[ERROR] = val
			case "wrn":
				levelColor[WARN] = val
			case "msg":
				levelColor[MSG] = val
			case "inf":
				levelColor[INFO] = val
			case "dbg":
				levelColor[DEBUG] = val
			}
		}
	}
}
