FROM uwspkg/build

LABEL mantainer="Jeremías Casteglione <jeremias.tincan@gmail.com>"
LABEL version="2021.02.14"

USER uws:uws
WORKDIR /home/uws

ENV USER uws
ENV HOME /home/uws

COPY --chown=uws:uws ./utils /usr/local/bin

RUN mkdir -vp ./build ./src
COPY --chown=uws:uws ./utils /usr/local/bin
COPY --chown=uws:uws ./files /home/uws/src

RUN mkpkg.sh

CMD exec /usr/local/bin/uws-login.sh
