// Copyright (c) Jeremías Casteglione <jeremias@talkingpts.org>
// See LICENSE file.

package uwspkg

import (
	"testing"

	_ "uwspkg/_testing/setup"

	"uwspkg/config"
)

func TestConfigDefaults(t *testing.T) {
	if len(config.ConfigFiles) != 1 {
		t.Fatalf("number of config files: got '%d' - expect '%d'", len(config.ConfigFiles), 1)
	}
}

func TestPackage(t *testing.T) {
}
