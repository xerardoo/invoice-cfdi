package excel_test

import (
	"github.com/peteole/testdata-loader"
	"github.com/stretchr/testify/assert"
	repository "github.com/xerardoo/invoice-cfdi-validator/internal/validator/repository/excel"
	"github.com/xerardoo/invoice-cfdi-validator/pkg/datasource"
	"path/filepath"
	"testing"
)

func connectUsoCFDIExcel() datasource.FileExcel {
	catalogSatFilePath := filepath.Join(testdataloader.GetBasePath(), "testdata/catCFDI_V_33_code_challenge.xls")
	fileExcel := datasource.NewFileExcel(catalogSatFilePath)
	err := fileExcel.Connect()
	if err != nil {
		panic(err)
	}
	return fileExcel
}

func TestUsoCFDI_GetAll(t *testing.T) {
	fileExcel := connectUsoCFDIExcel()
	usoCfdiRepo := repository.NewUsoCFDIExcelRepo(fileExcel)
	usoCfdiRepo = usoCfdiRepo.SetSheet("c_UsoCFDI", 1, []int{0, 2, 3}, 6, 10)

	usosCfdi, err := usoCfdiRepo.GetAll()

	assert.Equal(t, 5, len(usosCfdi))
	assert.Nil(t, err)
}

func TestUsoCFDI_GetByClave(t *testing.T) {
	fileExcel := connectUsoCFDIExcel()
	usosCfdiRepo := repository.NewUsoCFDIExcelRepo(fileExcel)
	usosCfdiRepo = usosCfdiRepo.SetSheet("c_UsoCFDI", 1, []int{0, 2, 3}, 6, 27)
	usoCfdiPorDefinir := "P01"

	usoCfdi, err := usosCfdiRepo.GetByClave(usoCfdiPorDefinir)

	assert.Nil(t, err)
	assert.Equal(t, usoCfdiPorDefinir, usoCfdi.Clave)
	assert.Equal(t, true, usoCfdi.IsFisica)
	assert.Equal(t, true, usoCfdi.IsMoral)
}
