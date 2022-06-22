package excel_test

import (
	"github.com/peteole/testdata-loader"
	"github.com/stretchr/testify/assert"
	repository "github.com/xerardoo/invoice-cfdi-validator/internal/validator/repository/excel"
	"github.com/xerardoo/invoice-cfdi-validator/pkg/datasource"
	"path/filepath"
	"testing"
)

func connectFormaPagoExcel() datasource.FileExcel {
	catalogSatFilePath := filepath.Join(testdataloader.GetBasePath(), "testdata/catCFDI_V_33_code_challenge.xls")
	fileExcel := datasource.NewFileExcel(catalogSatFilePath)
	err := fileExcel.Connect()
	if err != nil {
		panic(err)
	}
	return fileExcel
}

func TestFormaPago_GetAll(t *testing.T) {
	fileExcel := connectFormaPagoExcel()
	formaPagoRepo := repository.NewFormaPagoExcelRepo(fileExcel)
	formaPagoRepo = formaPagoRepo.SetSheet("c_FormaPago", 0, 0, 6, 27)

	formasPago, err := formaPagoRepo.GetAll()

	assert.Equal(t, 22, len(formasPago))
	assert.Nil(t, err)
}

func TestFormaPago_GetByClave(t *testing.T) {
	fileExcel := connectFormaPagoExcel()
	formaPagoRepo := repository.NewFormaPagoExcelRepo(fileExcel)
	formaPagoRepo = formaPagoRepo.SetSheet("c_FormaPago", 0, 0, 6, 6)
	formaPagoEfectivo := "01"

	formaPago, err := formaPagoRepo.GetByClave(formaPagoEfectivo)

	assert.Equal(t, formaPagoEfectivo, formaPago.Clave)
	assert.Nil(t, err)
}
