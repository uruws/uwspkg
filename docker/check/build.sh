#!/bin/sh
set -eu
exec docker build $@ --rm -t uwspkg/check ./docker/check
