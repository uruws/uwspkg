#!/bin/sh
exec docker run --rm --name uwspkg-golang-check \
	--hostname golang-check.uwspkg.local \
	-v ${PWD}/go:/go/src/uwspkg:ro \
	-e UWSPKG_LOG='debug' \
	-e UWSPKG_LOG_COLORS='off' \
	-u uws uwspkg/golang make check
