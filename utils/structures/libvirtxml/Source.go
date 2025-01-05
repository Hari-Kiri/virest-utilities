package libvirtxml

import (
	"encoding/xml"

	"github.com/Hari-Kiri/virest-utilities/utils/structures/virest"
	"libvirt.org/go/libvirt"
)

type Source struct {
	XMLName   xml.Name  `xml:"source"`
	Host      Host      `xml:"host"`
	Initiator Initiator `xml:"initiator"`
}

type Host struct {
	Name string `xml:"name,attr"`
	Port int    `xml:"port,attr"`
}

type Initiator struct {
	Iqn Iqn `xml:"iqn"`
}

type Iqn struct {
	Name string `xml:"name,attr"`
}

func (source *Source) Unmarshal(doc string) (virest.Error, bool) {
	errorUnmarshal := xml.Unmarshal([]byte(doc), source)
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

func (source *Source) Marshal() (string, virest.Error, bool) {
	doc, errorMarshal := xml.MarshalIndent(source, "", "  ")
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
