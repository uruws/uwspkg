#!/bin/sh
exec docker run -it --rm --network none --name uwspkg-base \
	--hostname base.uwspkg.local -u uws uwspkg/base $@
