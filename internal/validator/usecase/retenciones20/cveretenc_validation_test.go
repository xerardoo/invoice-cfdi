package retenciones20_test

import (
	"github.com/stretchr/testify/assert"
	usecase "github.com/xerardoo/invoice-cfdi-validator/internal/validator/usecase/retenciones20"
	"github.com/xerardoo/invoice-cfdi-validator/pkg/retenciones20"
	"testing"
)

func TestCveRetencUseCase_CveRetencExists(t *testing.T) {
	cfdi1 := retenciones20.NewRetenciones()
	cfdi1.CveRetenc = "25"
	cfdi1.DescRetenc = ""
	usecase1 := usecase.NewCveRetencUseCase(cfdi1)

	cfdi2 := retenciones20.NewRetenciones()
	cfdi2.CveRetenc = "25"
	cfdi2.DescRetenc = "20"
	usecase2 := usecase.NewCveRetencUseCase(cfdi2)

	var tests = []struct {
		caseName string
		input    usecase.CveRetencUseCase
		expected bool
	}{
		{
			"DescRetenc No debe existir",
			usecase1,
			false,
		}, {
			"DescRetenc debe existir",
			usecase2,
			true,
		},
	}

	for _, tt := range tests {
		t.Log(tt.caseName)

		isValid, _ := tt.input.CveRetencExists().Validate()
		assert.Equal(t, tt.expected, isValid)
	}
}
