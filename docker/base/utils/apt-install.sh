#!/bin/sh
set -eu
export DEBIAN_FRONTEND=noninteractive

apt-get clean
apt-get update

apt-get install -yy --no-install-recommends $@

apt-get clean
apt-get autoremove -yy --purge
rm -rf /var/lib/apt/lists/* \
	/var/cache/apt/archives/*.deb \
	/var/cache/apt/*cache.bin

exit 0
