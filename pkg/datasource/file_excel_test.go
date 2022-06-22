package datasource_test

import (
	"github.com/peteole/testdata-loader"
	"github.com/stretchr/testify/assert"
	"github.com/xerardoo/invoice-cfdi-validator/pkg/datasource"
	"path/filepath"
	"testing"
)

func TestNewFileExcel(t *testing.T) {
	filename := "myexcel.xls"

	newFileExcel := datasource.NewFileExcel(filename)

	assert.Equal(t, newFileExcel.Filename, filename)
}

func TestFileExcel_Connect(t *testing.T) {
	catalogSatFilePath := filepath.Join(testdataloader.GetBasePath(), "testdata/catCFDI_V_33_code_challenge.xls")
	newFileExcel := datasource.NewFileExcel(catalogSatFilePath)

	err := newFileExcel.Connect()

	assert.Nil(t, err)
}
