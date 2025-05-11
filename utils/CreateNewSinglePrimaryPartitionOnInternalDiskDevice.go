package utils

import (
	"strings"

	"github.com/Hari-Kiri/virest-utilities/utils/libguestfs"
	"github.com/Hari-Kiri/virest-utilities/utils/structures/virest"
	"libvirt.org/go/libvirt"
)

// Partition a local disk device with one primary partition.
// Parameters:
//
//   - diskDevicePath: disk device location (ex: /dev/sda or /home/user/image.qcow2).
//   - diskDeviceFormat: disk device format (ex: raw or qcow2).
//   - diskDevicePartitionTable: disk device partition table (ex: mbr or gpt).
func CreateNewSinglePrimaryPartitionOnInternalDiskDevice(diskDevicePath string, diskDeviceFormat string, diskDevicePartitionTable string) (string, virest.Error, bool) {
	guestfs, errorCreateLibguestfsHandle := libguestfs.Create()
	if errorCreateLibguestfsHandle != nil {
		return "", virest.Error{Error: libvirt.Error{
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
		return "", virest.Error{Error: libvirt.Error{
			Code:    libvirt.ERR_INTERNAL_ERROR,
			Domain:  libvirt.FROM_STORAGE,
			Message: errorAddDrive.Error(),
			Level:   libvirt.ERR_ERROR,
		}}, true
	}

	// run the libguestfs back-end
	if errorLaunchGuestfs := guestfs.Launch(); errorLaunchGuestfs != nil {
		return "", virest.Error{Error: libvirt.Error{
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
		return "", virest.Error{Error: libvirt.Error{
			Code:    libvirt.ERR_INTERNAL_ERROR,
			Domain:  libvirt.FROM_STORAGE,
			Message: errorGetListOfDevices.Error(),
			Level:   libvirt.ERR_ERROR,
		}}, true
	}
	if len(devices) > 1 {
		return "", virest.Error{Error: libvirt.Error{
			Code:    libvirt.ERR_INTERNAL_ERROR,
			Domain:  libvirt.FROM_STORAGE,
			Message: "expected a single device",
			Level:   libvirt.ERR_ERROR,
		}}, true
	}

	// partition the disk as one single partition
	errorPartitioningDisk := guestfs.Part_disk(devices[0], diskDevicePartitionTable)
	if errorPartitioningDisk != nil {
		return "", virest.Error{Error: libvirt.Error{
			Code:    libvirt.ERR_INTERNAL_ERROR,
			Domain:  libvirt.FROM_STORAGE,
			Message: errorPartitioningDisk.Error(),
			Level:   libvirt.ERR_ERROR,
		}}, true
	}

	// get the list of partitions
	// we expect a single element, which is the partition we have just created
	partitions, errorGetListOfDiskPartition := guestfs.List_partitions()
	if errorGetListOfDiskPartition != nil {
		return "", virest.Error{Error: libvirt.Error{
			Code:    libvirt.ERR_INTERNAL_ERROR,
			Domain:  libvirt.FROM_STORAGE,
			Message: errorGetListOfDiskPartition.Error(),
			Level:   libvirt.ERR_ERROR,
		}}, true
	}
	if len(partitions) > 1 {
		var partitionsString strings.Builder
		for i := 0; i < len(partitions); i++ {
			partitionsString.WriteString(partitions[i])
		}
		return partitionsString.String(), virest.Error{Error: libvirt.Error{
			Code:    libvirt.ERR_INTERNAL_ERROR,
			Domain:  libvirt.FROM_STORAGE,
			Message: "expected a single partition from list-partitions",
			Level:   libvirt.ERR_ERROR,
		}}, true
	}
	if len(partitions) < 1 {
		return "", virest.Error{Error: libvirt.Error{
			Code:    libvirt.ERR_INTERNAL_ERROR,
			Domain:  libvirt.FROM_STORAGE,
			Message: "failed create new single primary partition on internal disk device",
			Level:   libvirt.ERR_ERROR,
		}}, true
	}

	return partitions[0], virest.Error{}, false
}
