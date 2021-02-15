#!/bin/sh
set -eu
exec docker build $@ --rm -t uwspkg/pkg ./base/pkg
