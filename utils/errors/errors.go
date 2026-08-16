package errors

import (
	stderrors "errors"
	"fmt"

	"github.com/Hari-Kiri/virest-utilities/utils/structures/virest"
	"libvirt.org/go/libvirt"
)

// Wrap annotates err with an operation name. Nil input stays nil.
func Wrap(op string, err error) error {
	if err == nil {
		return nil
	}
	return fmt.Errorf("%s: %w", op, err)
}

// AsLibvirtError extracts a libvirt.Error from err when present.
func AsLibvirtError(err error) (libvirt.Error, bool) {
	var lv libvirt.Error
	if stderrors.As(err, &lv) {
		return lv, true
	}
	return libvirt.Error{}, false
}

// Code returns the libvirt error number when err unwraps to libvirt.Error.
func Code(err error) (libvirt.ErrorNumber, bool) {
	lv, ok := AsLibvirtError(err)
	if !ok {
		return 0, false
	}
	return lv.Code, true
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
