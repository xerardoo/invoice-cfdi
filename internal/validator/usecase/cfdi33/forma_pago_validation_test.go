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

func connectCatalogExcel() datasource.FileExcel {
	catalogSatFilePath := filepath.Join(testdataloader.GetBasePath(), "testdata/catCFDI_V_33_code_challenge.xls")
	fileExcel := datasource.NewFileExcel(catalogSatFilePath)
	err := fileExcel.Connect()
	if err != nil {
		panic(err)
	}
	return fileExcel
}

func TestFormaExist(t *testing.T) {
	fileExcel := connectCatalogExcel()
	cfdi1 := cfdi.NewCFDI()
	cfdi1.FormaPago = "01"
	cfdi2 := cfdi.NewCFDI()
	cfdi2.FormaPago = "100"

	formaPagoUseCase1 := usecase.NewFormaPagoUseCase(fileExcel, cfdi1)
	formaPagoUseCase2 := usecase.NewFormaPagoUseCase(fileExcel, cfdi2)

	var tests = []struct {
		caseName string
		input    validator.AttributeValidation
		expected bool
	}{
		{
			"Fecha valida",
			formaPagoUseCase1.FormaExists(),
			true,
		}, {
			"Fecha No valida",
			formaPagoUseCase2.FormaExists(),
			false,
		},
	}

	for _, tt := range tests {
		t.Log(tt.caseName)

		isValid, _ := tt.input.Validate()
		assert.Equal(t, tt.expected, isValid)
	}
}
