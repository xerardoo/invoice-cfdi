package cfdi33

import (
	"encoding/xml"
	"github.com/xerardoo/invoice-cfdi-validator/pkg/validator"
)

type CFDI struct {
	Metadata
	Emisor    Emisor
	Receptor  Receptor
	Conceptos Conceptos
	Impuestos Impuestos
}

type Metadata struct {
	XMLName           xml.Name `xml:"Comprobante"`
	Certificado       string   `xml:"Certificado,attr"`
	CondicionesDePago string   `xml:"CondicionesDePago,attr"`
	Fecha             string   `xml:"Fecha,attr"`
	Folio             string   `xml:"Folio,attr"`
	FormaPago         string   `xml:"FormaPago,attr"`
	LugarExpedicion   string   `xml:"LugarExpedicion,attr"`
	MetodoPago        string   `xml:"MetodoPago,attr"`
	Moneda            string   `xml:"Moneda,attr"`
	NoCertificado     string   `xml:"NoCertificado,attr"`
	Sello             string   `xml:"Sello,attr"`
	Serie             string   `xml:"Serie,attr"`
	SubTotal          string   `xml:"SubTotal,attr"`
	TipoCambio        string   `xml:"TipoCambio,attr"`
	TipoDeComprobante string   `xml:"TipoDeComprobante,attr"`
	Total             string   `xml:"Total,attr"`
	Version           string   `xml:"Version,attr"`
	Cfdi              string   `xml:"cfdi,attr"`
	Xsi               string   `xml:"xsi,attr"`
	SchemaLocation    string   `xml:"schemaLocation,attr"`
}

func NewCFDI() CFDI {
	return CFDI{}
}

func (cfdi CFDI) GetVersion() string {
	return cfdi.Version
}

func (cfdi CFDI) ValidateFields(validations []validator.IValidator) (errors []validator.IErrValidation) {
	for _, fieldValidation := range validations {
		if isValid, err := fieldValidation.Validate(); !isValid {
			errors = append(errors, err)
		}
	}
	return errors
}
