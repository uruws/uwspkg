#!/bin/sh
set -eu
doas install -v -d -m 0750 /etc/schroot/bootstrap-uwspkg
doas install -v -C -m 0640 ./etc/schroot/bootstrap-uwspkg/* /etc/schroot/bootstrap-uwspkg/
doas install -v -C -m 0640 ./etc/schroot/chroot.d/bootstrap-uwspkg.conf /etc/schroot/chroot.d/
#~ make
exit 0
