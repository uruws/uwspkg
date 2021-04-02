URI := https://download.freebsd.org/ftp/releases/VM-IMAGES/12.2-RELEASE/amd64/Latest/FreeBSD-12.2-RELEASE-amd64.qcow2.xz

hd0.qcow: fbsd.qcow2.xz
	@unxz -d -c fbsd.qcow2.xz >hd0.qcow

fbsd.qcow2.xz:
	@wget -c -O fbsd.qcow2.xz $(URI)

.PHONY: up
up:
	@kvm -m 1024 -name pkgng -hda ./hd0.qcow \
		-nic 'user,hostfwd=tcp:127.0.0.1:28128-:22'
