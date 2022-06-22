package retenciones20_test

import (
	"github.com/stretchr/testify/assert"
	usecase "github.com/xerardoo/invoice-cfdi-validator/internal/validator/usecase/retenciones20"
	"github.com/xerardoo/invoice-cfdi-validator/pkg/retenciones20"
	"testing"
)

func TestFechaUseCase_FechaExpFormat(t *testing.T) {
	cfdi1 := retenciones20.NewRetenciones()
	cfdi1.FechaExp = "2019-07-08T10:12:18"
	usecase1 := usecase.NewFechaExpUseCase(cfdi1)

	cfdi2 := retenciones20.NewRetenciones()
	cfdi2.FechaExp = "1999-01-01T10:12:18"
	usecase2 := usecase.NewFechaExpUseCase(cfdi2)

	var tests = []struct {
		caseName string
		input    usecase.FechaExpUseCase
		expected bool
	}{
		{
			"fecha es valida",
			usecase1,
			true,
		}, {
			"fecha No valida",
			usecase2,
			false,
		},
	}

	for _, tt := range tests {
		t.Log(tt.caseName)

		isValid, _ := tt.input.FechaExpFormat().Validate()
		assert.Equal(t, tt.expected, isValid)
	}
}
