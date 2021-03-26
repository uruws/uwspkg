#!/bin/sh
exec docker run -it --rm --name uwspkg-golang \
	--hostname golang.uwspkg.local \
	-v ${PWD}/base:/usr/local/src/base:ro \
	-v ${PWD}/devel:/usr/local/src/devel:ro \
	-v ${PWD}/go:/go/src/uwspkg \
	-e UWSPKG_LOG='debug' \
	-e UWSPKG_LOG_COLORS='auto' \
	-u uws uwspkg/golang $@
