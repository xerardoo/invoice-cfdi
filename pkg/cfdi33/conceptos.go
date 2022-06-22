package cfdi33

import "encoding/xml"

type Conceptos struct {
	XMLName   xml.Name `xml:"cfdi:Conceptos"`
	Conceptos []Concepto
}

type Concepto struct {
	XMLName          xml.Name `xml:"cfdi:Concepto"`
	NoIdentificacion string   `xml:"NoIdentificacion,attr"`
	ClaveProdServ    string   `xml:"ClaveProdServ,attr"`
	Cantidad         string   `xml:"Cantidad,attr"`
	ClaveUnidad      string   `xml:"ClaveUnidad,attr"`
	Unidad           string   `xml:"Unidad,attr"`
	Descripcion      string   `xml:"Descripcion,attr"`
	ValorUnitario    string   `xml:"ValorUnitario,attr"`
	Importe          string   `xml:"Importe,attr"`
	Descuento        string   `xml:"Descuento,attr"`
	Impuestos        Impuestos
}
