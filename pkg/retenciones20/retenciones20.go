package retenciones20

import "encoding/xml"

type Retenciones struct {
	XMLName        xml.Name `xml:"Retenciones"`
	SchemaLocation string   `xml:"schemaLocation,attr"`
	Version        string   `xml:"Version,attr"`
	FolioInt       string   `xml:"FolioInt,attr"`
	Sello          string   `xml:"Sello,attr"`
	NoCertificado  string   `xml:"NoCertificado,attr"`
	Certificado    string   `xml:"Certificado,attr"`
	FechaExp       string   `xml:"FechaExp,attr"`
	LugarExpRetenc string   `xml:"LugarExpRetenc,attr"`
	CveRetenc      string   `xml:"CveRetenc,attr"`
	DescRetenc     string   `xml:"DescRetenc,attr"`
	Retenciones    string   `xml:"retenciones,attr"`
	Xsi            string   `xml:"xsi,attr"`
	Emisor
	Receptor
	Periodo
	Totales struct {
		MontoTotOperacion string `xml:"MontoTotOperacion,attr"`
		MontoTotGrav      string `xml:"MontoTotGrav,attr"`
		MontoTotExent     string `xml:"MontoTotExent,attr"`
		MontoTotRet       string `xml:"MontoTotRet,attr"`
		ImpRetenidos      []struct {
			Text        string `xml:",chardata"`
			BaseRet     string `xml:"BaseRet,attr"`
			ImpuestoRet string `xml:"ImpuestoRet,attr"`
			MontoRet    string `xml:"MontoRet,attr"`
			TipoPagoRet string `xml:"TipoPagoRet,attr"`
		} `xml:"ImpRetenidos"`
	} `xml:"Totales"`
	Complemento struct {
		TimbreFiscalDigital struct {
			SchemaLocation   string `xml:"schemaLocation,attr"`
			Version          string `xml:"Version,attr"`
			UUID             string `xml:"UUID,attr"`
			FechaTimbrado    string `xml:"FechaTimbrado,attr"`
			RfcProvCertif    string `xml:"RfcProvCertif,attr"`
			SelloCFD         string `xml:"SelloCFD,attr"`
			NoCertificadoSAT string `xml:"NoCertificadoSAT,attr"`
			SelloSAT         string `xml:"SelloSAT,attr"`
			Tfd              string `xml:"tfd,attr"`
			Xsi              string `xml:"xsi,attr"`
		} `xml:"TimbreFiscalDigital"`
	} `xml:"Complemento"`
}

type Emisor struct {
	XMLName        xml.Name `xml:"Emisor"`
	RfcE           string   `xml:"RfcE,attr"`
	NomDenRazSocE  string   `xml:"NomDenRazSocE,attr"`
	RegimenFiscalE string   `xml:"RegimenFiscalE,attr"`
}

type Receptor struct {
	XMLName       xml.Name `xml:"Receptor"`
	NacionalidadR string   `xml:"NacionalidadR,attr"`
	Nacional      struct {
		Text             string `xml:",chardata"`
		RfcR             string `xml:"RfcR,attr"`
		NomDenRazSocR    string `xml:"NomDenRazSocR,attr"`
		DomicilioFiscalR string `xml:"DomicilioFiscalR,attr"`
	} `xml:"Nacional"`
}

type Periodo struct {
	XMLName   xml.Name `xml:"Periodo"`
	Text      string   `xml:",chardata"`
	MesIni    string   `xml:"MesIni,attr"`
	MesFin    string   `xml:"MesFin,attr"`
	Ejercicio string   `xml:"Ejercicio,attr"`
}

func (cfdi Retenciones) GetVersion() string {
	return cfdi.Version
}

func NewRetenciones() Retenciones {
	return Retenciones{}
}
