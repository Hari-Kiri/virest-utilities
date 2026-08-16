package hypervisor

import (
	"errors"
	"fmt"

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
		return virest.Connection{}, fmt.Errorf("%s: %w", op, err)
	}
	if conn == nil {
		return virest.Connection{}, fmt.Errorf("%s: nil libvirt connection", op)
	}
	return virest.Connection{Connect: conn}, nil
}

// AsLibvirtError extracts a libvirt.Error from err when present.
func AsLibvirtError(err error) (libvirt.Error, bool) {
	var lv libvirt.Error
	if errors.As(err, &lv) {
		return lv, true
	}
	return libvirt.Error{}, false
}

// ToLegacy converts an idiomatic error into the historical (virest.Error, bool) pair.
// Non-libvirt errors are preserved as a libvirt.Error with Message set so callers
// that only inspect Message still see the failure.
func ToLegacy(err error) (virest.Error, bool) {
	if err == nil {
		return virest.Error{}, false
	}
	if lv, ok := AsLibvirtError(err); ok {
		return virest.Error{Error: lv}, true
	}
	return virest.Error{Error: libvirt.Error{
		Code:    libvirt.ERR_INTERNAL_ERROR,
		Domain:  libvirt.FROM_NONE,
		Message: err.Error(),
		Level:   libvirt.ERR_ERROR,
	}}, true
}
