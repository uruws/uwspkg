#!/bin/sh
exec docker run -it --rm --name uwspkg-build-devel \
	--hostname build-devel.uwspkg.local -u uws uwspkg/build $@
