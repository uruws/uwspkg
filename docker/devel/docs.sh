#!/bin/sh
exec docker run -it --rm --name uwspkg-docs \
	--hostname docs.uwspkg.local \
	-e PATH=/uws/sbin:/usr/local/bin:/usr/bin:/bin \
	--workdir /home/uws \
	-u uws uwspkg/devel /bin/bash
