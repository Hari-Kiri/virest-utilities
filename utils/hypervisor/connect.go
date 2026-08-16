package hypervisor

import (
	"fmt"

	viresterrors "github.com/Hari-Kiri/virest-utilities/utils/errors"
	"github.com/Hari-Kiri/virest-utilities/utils/structures/virest"
	"libvirt.org/go/libvirt"
)

// Connect opens a connection to the hypervisor at uri.
// See https://libvirt.org/uri.html for URI formats (including local UNIX sockets
// such as qemu:///system and remote transports such as qemu+ssh://...).
//
// Close the returned connection with Connection.Close when finished.
func Connect(uri string) (virest.Connection, error) {
	conn, err := libvirt.NewConnect(uri)
	return wrapConnect(conn, err, "connect")
}

// ConnectReadOnly opens a read-only connection to the hypervisor at uri.
func ConnectReadOnly(uri string) (virest.Connection, error) {
	conn, err := libvirt.NewConnectReadOnly(uri)
	return wrapConnect(conn, err, "connect read-only")
}

// ConnectWithAuth opens a connection, invoking auth when credentials are required.
func ConnectWithAuth(uri string, auth *libvirt.ConnectAuth, flags libvirt.ConnectFlags) (virest.Connection, error) {
	conn, err := libvirt.NewConnectWithAuth(uri, auth, flags)
	return wrapConnect(conn, err, "connect with auth")
}

// ConnectWithAuthDefault opens a connection using libvirt's default auth callback.
func ConnectWithAuthDefault(uri string, flags libvirt.ConnectFlags) (virest.Connection, error) {
	conn, err := libvirt.NewConnectWithAuthDefault(uri, flags)
	return wrapConnect(conn, err, "connect with auth default")
}

func wrapConnect(conn *libvirt.Connect, err error, op string) (virest.Connection, error) {
	if err != nil {
		return virest.Connection{}, viresterrors.Wrap(op, err)
	}
	if conn == nil {
		return virest.Connection{}, fmt.Errorf("%s: nil libvirt connection", op)
	}
	return virest.Connection{Connect: conn}, nil
}
