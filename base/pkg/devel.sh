#!/bin/sh
exec docker run -it --rm --name uwspkg-pkg-devel \
	--hostname pkg-devel.uwspkg.local -u uws uwspkg/pkg $@
