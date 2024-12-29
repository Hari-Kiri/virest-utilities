package libvirtxml

import (
	"encoding/xml"
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

func (storagepoolCapabilities *StoragepoolCapabilities) Unmarshal(doc string) error {
	return xml.Unmarshal([]byte(doc), storagepoolCapabilities)
}

func (storagepoolCapabilities *StoragepoolCapabilities) Marshal() (string, error) {
	doc, err := xml.MarshalIndent(storagepoolCapabilities, "", "  ")
	if err != nil {
		return "", err
	}
	return string(doc), nil
}
