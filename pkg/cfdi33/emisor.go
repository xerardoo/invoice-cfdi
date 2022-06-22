package cfdi33

import "encoding/xml"

type Emisor struct {
	XMLName       xml.Name `xml:"cfdi:Emisor"`
	Nombre        string   `xml:"Nombre,attr"`
	RegimenFiscal string   `xml:"RegimenFiscal,attr"`
	RFC           string   `xml:"Rfc,attr"`
}

