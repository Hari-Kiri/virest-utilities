package libvirtxml

import (
	"encoding/xml"

	"github.com/Hari-Kiri/virest-utilities/utils/structures/virest"
	"libvirt.org/go/libvirt"
	"libvirt.org/go/libvirtxml"
)

type Sources struct {
	XMLName xml.Name                       `xml:"sources"`
	Source  []libvirtxml.StoragePoolSource `xml:"source"`
}

func (sources *Sources) Unmarshal(doc string) (virest.Error, bool) {
	errorUnmarshal := xml.Unmarshal([]byte(doc), sources)
	if errorUnmarshal != nil {
		return virest.Error{Error: libvirt.Error{
			Code:    libvirt.ERR_XML_ERROR,
			Domain:  libvirt.FROM_XML,
			Message: errorUnmarshal.Error(),
			Level:   libvirt.ERR_ERROR,
		}}, true
	}

	return virest.Error{}, false
}

func (sources *Sources) Marshal() (string, virest.Error, bool) {
	doc, errorMarshal := xml.MarshalIndent(sources, "", "  ")
	if errorMarshal != nil {
		return "", virest.Error{Error: libvirt.Error{
			Code:    libvirt.ERR_XML_ERROR,
			Domain:  libvirt.FROM_XML,
			Message: errorMarshal.Error(),
			Level:   libvirt.ERR_ERROR,
		}}, true
	}

	return string(doc), virest.Error{}, false
}
