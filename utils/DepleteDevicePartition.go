package utils

import (
	"github.com/Hari-Kiri/virest-utilities/utils/libguestfs"
	"github.com/Hari-Kiri/virest-utilities/utils/structures/virest"
	"libvirt.org/go/libvirt"
)

// Delete all partition on internal disk device.
//
// Parameters:
//
//   - diskDevicePath: disk device location (ex: /dev/sda or /home/user/image.qcow2).
//   - diskDeviceFormat: disk device format (ex: raw or qcow2).
func DepleteDevicePartition(diskDevicePath string, diskDeviceFormat string) (virest.Error, bool) {
	guestfs, errorCreateLibguestfsHandle := libguestfs.Create()
	if errorCreateLibguestfsHandle != nil {
		return virest.Error{Error: libvirt.Error{
			Code:    libvirt.ERR_INTERNAL_ERROR,
			Domain:  libvirt.FROM_STORAGE,
			Message: errorCreateLibguestfsHandle.Error(),
			Level:   libvirt.ERR_ERROR,
		}}, true
	}
	defer guestfs.Close()

	// attach the disk image to libguestfs
	optargs := libguestfs.OptargsAdd_drive{
		Format_is_set: true,
		Format:        diskDeviceFormat,
	}
	if errorAddDrive := guestfs.Add_drive(diskDevicePath, &optargs); errorAddDrive != nil {
		return virest.Error{Error: libvirt.Error{
			Code:    libvirt.ERR_INTERNAL_ERROR,
			Domain:  libvirt.FROM_STORAGE,
			Message: errorAddDrive.Error(),
			Level:   libvirt.ERR_ERROR,
		}}, true
	}

	// run the libguestfs back-end
	if errorLaunchGuestfs := guestfs.Launch(); errorLaunchGuestfs != nil {
		return virest.Error{Error: libvirt.Error{
			Code:    libvirt.ERR_INTERNAL_ERROR,
			Domain:  libvirt.FROM_STORAGE,
			Message: errorLaunchGuestfs.Error(),
			Level:   libvirt.ERR_ERROR,
		}}, true
	}

	// get the list of devices
	// we only expect that this list should contain a single device
	devices, errorGetListOfDevices := guestfs.List_devices()
	if errorGetListOfDevices != nil {
		return virest.Error{Error: libvirt.Error{
			Code:    libvirt.ERR_INTERNAL_ERROR,
			Domain:  libvirt.FROM_STORAGE,
			Message: errorGetListOfDevices.Error(),
			Level:   libvirt.ERR_ERROR,
		}}, true
	}
	if len(devices) > 1 {
		return virest.Error{Error: libvirt.Error{
			Code:    libvirt.ERR_INTERNAL_ERROR,
			Domain:  libvirt.FROM_STORAGE,
			Message: "expected a single device",
			Level:   libvirt.ERR_ERROR,
		}}, true
	}

	// get the list of partitions
	partitions, errorGetListOfDiskPartition := guestfs.List_partitions()
	if errorGetListOfDiskPartition != nil {
		return virest.Error{Error: libvirt.Error{
			Code:    libvirt.ERR_INTERNAL_ERROR,
			Domain:  libvirt.FROM_STORAGE,
			Message: errorGetListOfDiskPartition.Error(),
			Level:   libvirt.ERR_ERROR,
		}}, true
	}

	// Clear partition signature then delete all partition inside device
	var (
		errorWipefs          error
		errorDeletePartition error
	)
	if len(partitions) > 1 {
		for i := 0; i < len(partitions); i++ {
			errorWipefs = guestfs.Wipefs(partitions[i])
			errorDeletePartition = guestfs.Part_del(devices[0], i)
		}
	}
	if len(partitions) > 0 {
		errorWipefs = guestfs.Wipefs(partitions[0])
		errorDeletePartition = guestfs.Part_del(devices[0], 0)
	}
	if errorWipefs != nil {
		return virest.Error{Error: libvirt.Error{
			Code:    libvirt.ERR_INTERNAL_ERROR,
			Domain:  libvirt.FROM_STORAGE,
			Message: errorWipefs.Error(),
			Level:   libvirt.ERR_ERROR,
		}}, true
	}
	if errorDeletePartition != nil {
		return virest.Error{Error: libvirt.Error{
			Code:    libvirt.ERR_INTERNAL_ERROR,
			Domain:  libvirt.FROM_STORAGE,
			Message: errorDeletePartition.Error(),
			Level:   libvirt.ERR_ERROR,
		}}, true
	}

	return virest.Error{}, false
}
