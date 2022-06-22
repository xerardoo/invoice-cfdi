package cfdi33_test

import (
	"github.com/stretchr/testify/assert"
	usecase "github.com/xerardoo/invoice-cfdi-validator/internal/validator/usecase/cfdi33"
	"github.com/xerardoo/invoice-cfdi-validator/pkg/cfdi33"
	"testing"
)

func TestFechaUseCase_FechaFormat(t *testing.T) {
	cfdi1 := cfdi33.NewCFDI()
	cfdi1.Fecha = "2019-07-08T10:12:18"
	cfdi2 := cfdi33.NewCFDI()
	cfdi2.Fecha = "1999-01-01T10:12:18"

	var tests = []struct {
		caseName string
		input    usecase.FechaUseCase
		expected bool
	}{
		{
			"fecha es valida",
			usecase.FechaUseCase{Cfdi: cfdi1},
			true,
		}, {
			"fecha No valida",
			usecase.FechaUseCase{Cfdi: cfdi2},
			false,
		},
	}

	for _, tt := range tests {
		t.Log(tt.caseName)

		isValid, _ := tt.input.FechaFormat().Validate()
		assert.Equal(t, tt.expected, isValid)
	}
}
