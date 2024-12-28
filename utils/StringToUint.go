package utils

import (
	"strconv"

	"github.com/Hari-Kiri/virest-utilities/utils/structures/virest"
	"libvirt.org/go/libvirt"
)

// Convert string to uint. The uint is an unsigned integer type that is at least 32
// bits in size. It is a distinct type, however, and not an alias for, say, uint32.
func StringToUint(stringNumber string) (uint, virest.Error, bool) {
	var (
		result          uint64
		errorConverting error
	)

	result, errorConverting = strconv.ParseUint(stringNumber, 10, 32)
	if errorConverting != nil {
		libvirtError := libvirt.Error{
			Code:    libvirt.ERR_INVALID_ARG,
			Domain:  libvirt.FROM_NET,
			Message: "'Inactive' value not number or not exist",
			Level:   libvirt.ERR_ERROR,
		}
		return uint(result), virest.Error{Error: libvirtError}, true
	}

	return uint(result), virest.Error{}, false
}

// Convert string to uint. The uint32 is the set of all unsigned 32-bit integers. Range: 0 through 4294967295.
func StringToUint32(stringNumber string) (uint32, virest.Error, bool) {
	var (
		result          uint64
		errorConverting error
	)

	result, errorConverting = strconv.ParseUint(stringNumber, 10, 32)
	if errorConverting != nil {
		libvirtError := libvirt.Error{
			Code:    libvirt.ERR_INVALID_ARG,
			Domain:  libvirt.FROM_NET,
			Message: "'Inactive' value not number or not exist",
			Level:   libvirt.ERR_ERROR,
		}
		return uint32(result), virest.Error{Error: libvirtError}, true
	}

	return uint32(result), virest.Error{}, false
}

// Convert string to uint. The uint64 is the set of all unsigned 64-bit integers. Range: 0 through 18446744073709551615.
func StringToUint64(stringNumber string) (uint64, virest.Error, bool) {
	var (
		result          uint64
		errorConverting error
	)

	result, errorConverting = strconv.ParseUint(stringNumber, 10, 32)
	if errorConverting != nil {
		libvirtError := libvirt.Error{
			Code:    libvirt.ERR_INVALID_ARG,
			Domain:  libvirt.FROM_NET,
			Message: "'Inactive' value not number or not exist",
			Level:   libvirt.ERR_ERROR,
		}
		return result, virest.Error{Error: libvirtError}, true
	}

	return result, virest.Error{}, false
}
