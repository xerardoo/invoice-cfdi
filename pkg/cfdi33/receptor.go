package cfdi33

import "encoding/xml"

type Receptor struct {
	XMLName xml.Name `xml:"Receptor"`
	Nombre  string   `xml:"Nombre,attr"`
	RFC     string   `xml:"Rfc,attr"`
	UsoCFDI string   `xml:"UsoCFDI,attr"`
}
