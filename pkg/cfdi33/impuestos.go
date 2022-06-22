package cfdi33

import "encoding/xml"

type Impuestos struct {
	XMLName                   xml.Name `xml:"cfdi:Impuestos"`
	Traslados                 Traslados
	Retenciones               Retenciones
	TotalImpuestosTrasladados string `xml:"TotalImpuestosTrasladados,attr"`
	TotalImpuestosRetenidos   string `xml:"TotalImpuestosRetenidos,attr"`
}

type Traslados struct {
	XMLName  xml.Name `xml:"cfdi:Traslados"`
	Traslado []Traslado
}

type Retenciones struct {
	XMLName   xml.Name `xml:"cfdi:Retenciones"`
	Retencion []Retencion
}

type Traslado struct {
	XMLName    xml.Name `xml:"cfdi:Traslado"`
	Base       string   `xml:"Base,attr"`
	Impuesto   string   `xml:"Impuesto,attr"`
	TipoFactor string   `xml:"TipoFactor,attr"`
	TasaOCuota string   `xml:"TasaOCuota,attr"`
	Importe    string   `xml:"Importe,attr"`
}

type Retencion struct {
	XMLName    xml.Name `xml:"cfdi:Retencion"`
	Base       string   `xml:"Base,attr"`
	Impuesto   string   `xml:"Impuesto,attr"`
	TipoFactor string   `xml:"TipoFactor,attr"`
	TasaOCuota string   `xml:"TasaOCuota,attr"`
	Importe    string   `xml:"Importe,attr"`
}
