# uwspkg

uws packaging system.

## Requirements.

* Install *doas* (and grant root access for yourself):

    # apt-get install doas
    # echo 'permit nopass keepenv setenv { PATH } you as root' >>/etc/doas.conf

* Install *schroot*:

    # apt-get install schroot

* Create OS user `uwsbuild` and add yourself to it:

    # addgroup uwsbuild
    # adduser you uwsbuild

You will probably (most surely) have to re-login in order for the new membership
to take effect.

## Development.

* Bootstrap FreeBSD pkgng:

    $ make bootstrap

* Setup local environment:

    $ make setup

## Build all packages:

    $ make all
