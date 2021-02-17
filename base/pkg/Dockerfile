FROM uwspkg/build

LABEL mantainer="Jeremías Casteglione <jeremias.tincan@gmail.com>"
LABEL version="2021.02.16-1"

USER root:root
WORKDIR /root

ENV USER root
ENV HOME /root

RUN /root/bin/apt-install.sh fakeroot
RUN chgrp uws /usr/local/bin \
	&& chmod g+w /usr/local/bin

USER uws:uws
WORKDIR /home/uws

ENV USER uws
ENV HOME /home/uws

RUN mkdir -vp ./build ./src
VOLUME /home/uws/build

COPY --chown=uws:uws ./files /home/uws/src

ENV UWSPKG_VERSION 20210216

COPY --chown=uws:uws ./utils /usr/local/bin
ENTRYPOINT exec /usr/local/bin/mkpkg.sh
