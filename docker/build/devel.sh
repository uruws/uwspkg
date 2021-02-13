#!/bin/sh
exec docker run -it --rm --name uwspkg-devel \
	--hostname devel.uwspkg.local -u uws uwspkg/build $@
