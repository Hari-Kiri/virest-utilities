package utils

import (
	viresterrors "github.com/Hari-Kiri/virest-utilities/utils/errors"
	"github.com/Hari-Kiri/virest-utilities/utils/hypervisor"
	"github.com/Hari-Kiri/virest-utilities/utils/structures/virest"
	"libvirt.org/go/libvirt"
)

// NewConnect opens a connection to the hypervisor at uri.
//
// Deprecated: prefer hypervisor.Connect for idiomatic (Connection, error) results
// and to avoid importing the parent utils package when you only need libvirt connect.
//
// URIs are documented at https://libvirt.org/uri.html
// Close() should be used to release the resources after the connection is no longer needed.
func NewConnect(uri string) (virest.Connection, virest.Error, bool) {
	conn, err := hypervisor.Connect(uri)
	ve, isErr := viresterrors.ToLegacy(err)
	return conn, ve, isErr
}

// NewConnectReadOnly opens a read-only connection to the hypervisor at uri.
//
// Deprecated: prefer hypervisor.ConnectReadOnly.
//
// URIs are documented at https://libvirt.org/uri.html
func NewConnectReadOnly(uri string) (virest.Connection, virest.Error, bool) {
	conn, err := hypervisor.ConnectReadOnly(uri)
	ve, isErr := viresterrors.ToLegacy(err)
	return conn, ve, isErr
}

// NewConnectWithAuth opens a connection, invoking auth when credentials are required.
//
// Deprecated: prefer hypervisor.ConnectWithAuth.
//
// URIs are documented at https://libvirt.org/uri.html
func NewConnectWithAuth(uri string, auth *libvirt.ConnectAuth, flags libvirt.ConnectFlags) (virest.Connection, virest.Error, bool) {
	conn, err := hypervisor.ConnectWithAuth(uri, auth, flags)
	ve, isErr := viresterrors.ToLegacy(err)
	return conn, ve, isErr
}

// NewConnectWithAuthDefault opens a connection using libvirt's default auth callback.
//
// Deprecated: prefer hypervisor.ConnectWithAuthDefault.
//
// URIs are documented at https://libvirt.org/uri.html
func NewConnectWithAuthDefault(uri string, flags libvirt.ConnectFlags) (virest.Connection, virest.Error, bool) {
	conn, err := hypervisor.ConnectWithAuthDefault(uri, flags)
	ve, isErr := viresterrors.ToLegacy(err)
	return conn, ve, isErr
}
