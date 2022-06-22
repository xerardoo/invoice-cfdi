package retenciones20_test

import (
	"github.com/stretchr/testify/assert"
	usecase "github.com/xerardoo/invoice-cfdi-validator/internal/validator/usecase/retenciones20"
	"github.com/xerardoo/invoice-cfdi-validator/pkg/retenciones20"
	"testing"
)

func TestMesIniValidationUseCase_FechaFormat(t *testing.T) {
	reten1 := retenciones20.NewRetenciones()
	reten1.Periodo.MesIni = "01"
	reten1.Periodo.MesFin = "02"

	reten2 := retenciones20.NewRetenciones()
	reten2.Periodo.MesIni = "03"
	reten2.Periodo.MesFin = "02"

	var tests = []struct {
		caseName string
		input    usecase.MesIniUseCase
		expected bool
	}{
		{
			"fecha es valida",
			usecase.MesIniUseCase{Retenciones: reten1},
			true,
		}, {
			"fecha No valida",
			usecase.MesIniUseCase{Retenciones: reten2},
			false,
		},
	}

	for _, tt := range tests {
		t.Log(tt.caseName)

		isValid, _ := tt.input.MesIniValidation().Validate()
		assert.Equal(t, tt.expected, isValid)
	}
}
