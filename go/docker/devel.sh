#!/bin/sh
exec docker run -it --rm --name uwspkg-golang \
	--hostname golang.uwspkg.local -u uws uwspkg/golang $@
