package libvirtxml

import (
	"encoding/xml"
	"fmt"

	"github.com/Hari-Kiri/virest-utilities/utils/structures/virest"
	"libvirt.org/go/libvirt"
)

// Storage Pool Capabilities xml not found inside https://pkg.go.dev/libvirt.org/go/libvirtxml.
type StoragepoolCapabilities struct {
	XMLName xml.Name `xml:"storagepoolCapabilities"`
	Pool    []Pool   `xml:"pool"`
}

type Pool struct {
	Type        string      `xml:"type,attr"`
	Supported   string      `xml:"supported,attr"`
	VolOptions  VolOptions  `xml:"volOptions"`
	PoolOptions PoolOptions `xml:"poolOptions"`
}

type VolOptions struct {
	DefaultFormat DefaultFormat `xml:"defaultFormat"`
	Enum          Enum          `xml:"enum"`
}

type PoolOptions struct {
	DefaultFormat DefaultFormat `xml:"defaultFormat"`
	Enum          Enum          `xml:"enum"`
}

type DefaultFormat struct {
	Type string `xml:"type,attr"`
}

type Enum struct {
	Name  string   `xml:"name,attr"`
	Value []string `xml:"value"`
}

func (storagepoolCapabilities *StoragepoolCapabilities) Unmarshal(doc string) (virest.Error, bool) {
	errorUnmarshal := xml.Unmarshal([]byte(doc), storagepoolCapabilities)
	if errorUnmarshal != nil {
		return virest.Error{Error: libvirt.Error{
			Code:    libvirt.ERR_XML_ERROR,
			Domain:  libvirt.FROM_XML,
			Message: fmt.Sprintf("%s", errorUnmarshal),
			Level:   libvirt.ERR_ERROR,
		}}, true
	}

	return virest.Error{}, false
}

func (storagepoolCapabilities *StoragepoolCapabilities) Marshal() (string, virest.Error, bool) {
	doc, errorMarshal := xml.MarshalIndent(storagepoolCapabilities, "", "  ")
	if errorMarshal != nil {
		return "", virest.Error{Error: libvirt.Error{
			Code:    libvirt.ERR_XML_ERROR,
			Domain:  libvirt.FROM_XML,
			Message: fmt.Sprintf("%s", errorMarshal),
			Level:   libvirt.ERR_ERROR,
		}}, true
	}

	return string(doc), virest.Error{}, false
}
