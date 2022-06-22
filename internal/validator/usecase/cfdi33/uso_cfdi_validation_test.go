package cfdi33_test

import (
	"github.com/peteole/testdata-loader"
	"github.com/stretchr/testify/assert"
	usecase "github.com/xerardoo/invoice-cfdi-validator/internal/validator/usecase/cfdi33"
	cfdi "github.com/xerardoo/invoice-cfdi-validator/pkg/cfdi33"
	"github.com/xerardoo/invoice-cfdi-validator/pkg/datasource"
	"github.com/xerardoo/invoice-cfdi-validator/pkg/validator"
	"path/filepath"
	"testing"
)

func connectUsoCatalogExcel() datasource.FileExcel {
	catalogSatFilePath := filepath.Join(testdataloader.GetBasePath(), "testdata/catCFDI_V_33_code_challenge.xls")
	fileExcel := datasource.NewFileExcel(catalogSatFilePath)
	err := fileExcel.Connect()
	if err != nil {
		panic(err)
	}
	return fileExcel
}

func TestUsoCFDIExist(t *testing.T) {
	fileExcel := connectUsoCatalogExcel()
	cfdi1 := cfdi.NewCFDI()
	cfdi1.Receptor.RFC = "XXX010101XXX" // moral
	cfdi1.Receptor.UsoCFDI = "P01"
	cfdi2 := cfdi.NewCFDI()
	cfdi1.Receptor.RFC = "XXX0101010XXX" // fisica
	cfdi1.Receptor.UsoCFDI = "D01"

	usoCfdiUseCase1 := usecase.NewUsoCFDIUseCase(fileExcel, cfdi1)
	usoCfdiUseCase2 := usecase.NewUsoCFDIUseCase(fileExcel, cfdi2)

	var tests = []struct {
		caseName string
		input    validator.AttributeValidation
		expected bool
	}{
		{
			"Uso CFDI valido",
			usoCfdiUseCase1.UsoCfdiReceptor(),
			true,
		}, {
			"Uso CFDI No valido",
			usoCfdiUseCase2.UsoCfdiReceptor(),
			false,
		},
	}

	for _, tt := range tests {
		t.Log(tt.caseName)

		isValid, _ := tt.input.Validate()
		assert.Equal(t, tt.expected, isValid)
	}
}
