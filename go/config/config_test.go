// Copyright (c) Jeremías Casteglione <jeremias@talkingpts.org>
// See LICENSE file.

package config

import (
	"testing"
)

func TestDefaults(t *testing.T) {
	if len(ConfigFiles) != 3 {
		t.Fatalf("number of config files: got '%d' - expect '%d'", len(ConfigFiles), 3)
	}
}
