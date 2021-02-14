FROM uwspkg/base

LABEL mantainer="Jeremías Casteglione <jeremias.tincan@gmail.com>"
LABEL version="2021.02.13"

USER root:root
WORKDIR /root

ENV USER root
ENV HOME /root

ENV PKG 1.16.2

RUN mkdir -vp -m 0750 /home/uws/src
RUN wget -O pkgng.tgz https://github.com/freebsd/pkg/archive/${PKG}.tar.gz \
	&& tar -C /home/uws/src -xzf pkgng.tgz \
	&& chown -R uws:uws /home/uws/src

RUN /root/bin/apt-install.sh clang

RUN /root/bin/apt-install.sh autoconf automake libtool pkgconf make \
	freebsd-buildutils libarchive-dev libbz2-dev liblzma-dev libssl-dev

RUN mkdir -vp /uws && chown -vR uws:uws /uws

USER uws:uws
WORKDIR /home/uws

ENV USER uws
ENV HOME /home/uws

WORKDIR /home/uws/src/pkg-${PKG}
RUN CC=clang ./configure --prefix=/uws \
	&& make check CC=clang CFLAGS='-D_XOPEN_SOURCE=700'

RUN mkdir -vp /uws/etc \
	&& make install

CMD exec /usr/local/bin/uws-login.sh
