package virest

import "libvirt.org/go/libvirt"

type Result struct {
	Connection Connection
	Error      Error
}

type Connection struct {
	*libvirt.Connect
}

type Error struct {
	libvirt.Error
}
