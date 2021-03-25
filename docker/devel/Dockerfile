FROM uwspkg/build

LABEL mantainer="Jeremías Casteglione <jeremias.tincan@gmail.com>"
LABEL version="2021.02.18"

USER root:root
WORKDIR /root

ENV USER root
ENV HOME /root

RUN /root/bin/apt-install.sh sudo \
	&& printf 'uws\tALL=(ALL) NOPASSWD:ALL' >/etc/sudoers.d/uws \
	&& chmod -v 0440 /etc/sudoers.d/uws

RUN /root/bin/apt-install.sh man less vim-tiny fakeroot

RUN chgrp -v uws /usr/local/bin \
	&& chmod g+ws /usr/local/bin

RUN mkdir -vp /opt/src/uwspkg
VOLUME /opt/src/uwspkg

USER uws:uws
WORKDIR /home/uws

ENV USER uws
ENV HOME /home/uws

WORKDIR /opt/src/uwspkg
CMD exec /usr/local/bin/uws-login.sh
