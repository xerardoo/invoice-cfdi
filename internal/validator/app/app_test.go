package app_test

import (
	"encoding/xml"
	"fmt"
	"github.com/antchfx/xmlquery"
	"github.com/stretchr/testify/assert"
	"path/filepath"

	"github.com/peteole/testdata-loader"

	"github.com/spf13/viper"
	"github.com/xerardoo/invoice-cfdi-validator/internal/validator/app"
	"github.com/xerardoo/invoice-cfdi-validator/pkg/cfdi33"
	"io/ioutil"
	"os"
	"testing"
)

var (
	cfdi33FilePath = filepath.Join(testdataloader.GetBasePath(), "testdata/cfdi33.xml")
)

func TestRun(t *testing.T) {
	configFilePath := filepath.Join(testdataloader.GetBasePath(), "testdata/test.config.yaml")
	catalogSatFilePath := filepath.Join(testdataloader.GetBasePath(), "testdata/catCFDI_V_33_code_challenge.xls")

	viper.SetConfigFile(configFilePath)
	err := viper.ReadInConfig()
	if err != nil {
		t.Error(err)
	}
	viper.Set("sat_catalog_v33.datasource", catalogSatFilePath)

	err = app.Run(app.Args{
		Filepath: cfdi33FilePath,
	})

	assert.Nil(t, err)
}

func BenchmarkRun(b *testing.B) {
	configFilePath := filepath.Join(testdataloader.GetBasePath(), "testdata/test.config.yaml")
	catalogSatFilePath := filepath.Join(testdataloader.GetBasePath(), "testdata/catCFDI_V_33_code_challenge.xls")

	viper.SetConfigFile(configFilePath)
	err := viper.ReadInConfig()
	if err != nil {
		b.Error(err)
	}
	viper.Set("sat_catalog_v33.datasource", catalogSatFilePath)

	for i := 0; i < b.N; i++ {
		err := app.Run(app.Args{
			Filepath: cfdi33FilePath,
		})
		if err != nil {
			b.Error(err)
		}
	}
}

func BenchmarkXmlQuery(b *testing.B) {
	for i := 0; i < b.N; i++ {
		xmlFile, err := os.Open(cfdi33FilePath)
		if err != nil {
			b.Error(err)
		}
		defer xmlFile.Close()

		doc, err := xmlquery.Parse(xmlFile)
		if err != nil {
			b.Error(err)
		}

		xmlMeta := xmlquery.FindOne(doc, "/*") // get first node
		_ = fmt.Sprintf("%s:%s", xmlMeta.Prefix, xmlMeta.Data)
		_ = xmlMeta.SelectAttr("Version")
	}
}

func BenchmarkXmlMarshal(b *testing.B) {
	for i := 0; i < b.N; i++ {
		xmlFile, err := os.Open(cfdi33FilePath)
		if err != nil {
			b.Error(err)
		}
		defer xmlFile.Close()

		bts, _ := ioutil.ReadAll(xmlFile)

		var cfdi cfdi33.CFDI
		err = xml.Unmarshal(bts, &cfdi)
		if err != nil {
			b.Error(err)
		}
		_ = cfdi.GetVersion()
	}
}
