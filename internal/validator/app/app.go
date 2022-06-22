package app

import (
	"bytes"
	"encoding/json"
	"encoding/xml"
	"errors"
	"fmt"
	"github.com/antchfx/xmlquery"
	"github.com/spf13/viper"
	"github.com/xerardoo/invoice-cfdi-validator/internal/validator/usecase"
	cfdi33UC "github.com/xerardoo/invoice-cfdi-validator/internal/validator/usecase/cfdi33"
	ret20UC "github.com/xerardoo/invoice-cfdi-validator/internal/validator/usecase/retenciones20"
	"github.com/xerardoo/invoice-cfdi-validator/pkg/cfdi33"
	"github.com/xerardoo/invoice-cfdi-validator/pkg/datasource"
	"github.com/xerardoo/invoice-cfdi-validator/pkg/retenciones20"
	"github.com/xerardoo/invoice-cfdi-validator/pkg/validator"
	"io"
	"io/ioutil"
	"os"
	"strings"
)

var (
	MAX_SIZE_FILE_ALLOWED     int64 = 1048576 // 1Mb
	errMaxFileSize                  = errors.New("archivo demasiado grande. el maximo permitido es de 1 MB.")
	errOnlyFiles                    = errors.New("solo se permite archivos, carpetas no estan soportadas.")
	errOnlyXml                      = errors.New("solo archivos xml")
	errXmlNoSupported               = errors.New("XML no permitido.")
	errCfdiVersionNoSupported       = errors.New("version cfdi no soportada.")
)

type Args struct {
	Filepath string
}

func ReadConfig() {
	viper.SetConfigFile(`config.yaml`)
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
}

func NewAppArgs() Args {
	return Args{}
}

func Run(args Args) (err error) {
	xmlFile, err := os.Open(args.Filepath)
	if err != nil {
		err = fmt.Errorf("asui %s", err.Error())
		return
	}
	defer xmlFile.Close()

	// xml file validations
	xmlInfo, err := xmlFile.Stat()
	if err != nil {
		return
	}
	if xmlInfo.Size() >= MAX_SIZE_FILE_ALLOWED {
		err = errMaxFileSize
		return
	}
	if xmlInfo.IsDir() {
		err = errOnlyFiles
		return
	}
	if !strings.HasSuffix(xmlInfo.Name(), ".xml") {
		err = errOnlyXml
		return
	}

	// duplicates reader
	var buf bytes.Buffer
	tee := io.TeeReader(xmlFile, &buf)

	docXml, err := xmlquery.Parse(tee)
	if err != nil {
		panic(err)
	}
	xmlBytes, err := ioutil.ReadAll(&buf)
	if err != nil {
		return err
	}

	// determine kind and version of document
	xmlMeta := xmlquery.FindOne(docXml, "/*") // get first node
	xmlType := fmt.Sprintf("%s:%s", xmlMeta.Prefix, xmlMeta.Data)
	xmlVersion := xmlMeta.SelectAttr("Version")

	switch xmlType {
	case "cfdi:Comprobante":
		if xmlVersion != "3.3" {
			err = errCfdiVersionNoSupported
			return err
		}
		cfdi := cfdi33.NewCFDI()
		err = xml.Unmarshal(xmlBytes, &cfdi)
		if err != nil {
			return err
		}

		// open catalog file
		satCatalogDatasource := viper.GetString(`sat_catalog_v33.datasource`)
		fileExcel := datasource.NewFileExcel(satCatalogDatasource)
		err = fileExcel.Connect()
		if err != nil {
			return
		}

		formaPagoUseCase := cfdi33UC.NewFormaPagoUseCase(fileExcel, cfdi)
		usoCfdiUseCase := cfdi33UC.NewUsoCFDIUseCase(fileExcel, cfdi)
		fechaUseCase := cfdi33UC.NewFechaUseCase(cfdi)

		// validatios to do
		var validations = []validator.AttributeValidation{
			formaPagoUseCase.FormaExists(),
			usoCfdiUseCase.UsoCfdiReceptor(),
			fechaUseCase.FechaFormat(),
		}

		// exec validations
		runValidations(cfdi.GetVersion(), validations)
		break
	case "retenciones:Retenciones":
		if xmlVersion != "2.0" {
			err = errCfdiVersionNoSupported
			return err
		}
		retenciones := retenciones20.NewRetenciones()
		err = xml.Unmarshal(xmlBytes, &retenciones)
		if err != nil {
			return err
		}

		cveRetencUseCase := ret20UC.NewCveRetencUseCase(retenciones)
		fechaExpUseCase := ret20UC.NewFechaExpUseCase(retenciones)
		mesIniUseCase := ret20UC.NewMesIniUseCase(retenciones)

		// validatios to do
		var validations = []validator.AttributeValidation{
			cveRetencUseCase.CveRetencExists(),
			fechaExpUseCase.FechaExpFormat(),
			mesIniUseCase.MesIniValidation(),
		}

		// exec validations
		runValidations(retenciones.GetVersion(), validations)
		break
	default:
		err = errXmlNoSupported
		return err
	}
	return
}

func runValidations(cfdiVersion string, validations []validator.AttributeValidation) {
	var validationResult []*usecase.CFDIValidationMessage
	for _, attribute := range validations {
		isValid, _ := attribute.Validate()

		var newMessage = usecase.NewValidationMessage()
		newMessage.SetVersion(cfdiVersion)
		newMessage.SetCode(attribute.GetCode())
		newMessage.SetIsValid(isValid)
		if isValid {
			newMessage.SetMessage("OK")
		} else {
			newMessage.SetMessage(attribute.GetErrMessage())
		}
		validationResult = append(validationResult, newMessage)
	}

	result, _ := json.Marshal(&validationResult)
	fmt.Println(string(result))
}
