#!/bin/sh
exec docker run -it --rm --name uwspkg-golang \
	--hostname golang.uwspkg.local \
	-v ${PWD}/go:/go/src/uwspkg \
	-u uws uwspkg/golang $@
