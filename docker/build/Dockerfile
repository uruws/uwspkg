FROM uwspkg/base

LABEL mantainer="Jeremías Casteglione <jeremias.tincan@gmail.com>"
LABEL version="2021.02.13"

USER uws:uws
WORKDIR /home/uws

ENV USER uws
ENV HOME /home/uws

ENV PKG 1.16.2

RUN mkdir -vp -m 0750 /home/uws/src
WORKDIR /home/uws/src

RUN wget -O pkg-${PKG}.tgz https://github.com/freebsd/pkg/archive/${PKG}.tar.gz

USER root:root
WORKDIR /root

ENV USER root
ENV HOME /root

RUN /root/bin/apt-install.sh clang

RUN /root/bin/apt-install.sh autoconf automake libtool pkgconf make \
	freebsd-buildutils libarchive-dev libbz2-dev liblzma-dev libssl-dev

RUN mkdir -vp /uws && chown -vR uws:uws /uws

USER uws:uws
WORKDIR /home/uws

ENV USER uws
ENV HOME /home/uws

WORKDIR /home/uws/src
RUN tar -xzf pkg-${PKG}.tgz \
	&& cd pkg-${PKG} \
	&& CC=clang ./configure --prefix=/uws \
	&& make check CC=clang CFLAGS='-D_XOPEN_SOURCE=700' \
	&& mkdir -vp -m 0755 /uws/etc \
	&& make install \
	&& cd .. \
	&& rm -rf pkg-${PKG}

WORKDIR /home/uws
CMD exec /usr/local/bin/uws-login.sh
