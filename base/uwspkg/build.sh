#!/bin/sh
set -eu
exec docker build $@ --rm -t uwspkg/base:uwspkg ./base/uwspkg
