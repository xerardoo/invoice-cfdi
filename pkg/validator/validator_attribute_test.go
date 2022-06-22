package validator_test

import (
	"github.com/stretchr/testify/assert"
	"github.com/xerardoo/invoice-cfdi-validator/pkg/validator"
	"reflect"
	"runtime"
	"strconv"
	"testing"
)

func TestAttributeValidation_All(t *testing.T) {
	attributeValidation := validator.AttributeValidation{
		Name:       "Total",
		Value:      "100.00",
		ErrMessage: "Monto no valido.",
		Rule:       func(attributeValue string) bool { return false },
	}

	newAttrValidation := validator.NewAttributeValidation(
		attributeValidation.Name,
		attributeValidation.Value,
		attributeValidation.Code,
		attributeValidation.ErrMessage,
		attributeValidation.Rule,
	)

	assert.Equal(t, newAttrValidation.GetName(), attributeValidation.Name)
	assert.Equal(t, newAttrValidation.GetValue(), attributeValidation.Value)
	assert.Equal(t, newAttrValidation.GetCode(), attributeValidation.Code)
	assert.Equal(t, newAttrValidation.GetErrMessage(), attributeValidation.ErrMessage)

	ruleGot := runtime.FuncForPC(reflect.ValueOf(newAttrValidation.GetRule()).Pointer()).Name()
	ruleExpected := runtime.FuncForPC(reflect.ValueOf(attributeValidation.Rule).Pointer()).Name()
	assert.Equal(t, ruleGot, ruleExpected)
}

func TestAttributeValidation_Validate(t *testing.T) {
	var tests = []struct {
		caseName string
		input    validator.AttributeValidation
		expected bool
	}{
		{"no valid",
			validator.AttributeValidation{
				Name:       "Moneda",
				Value:      "KURKS",
				Code:       "MXN01",
				ErrMessage: "Moneda no valida.",
				Rule: func(attributeValue string) bool {
					if attributeValue != "MXN" {
						return false
					}
					return true
				},
			},
			false,
		},
		{"valid",
			validator.AttributeValidation{
				Name:       "Total",
				Value:      "999.99",
				Code:       "GHS500",
				ErrMessage: "Total no valido.",
				Rule: func(attributeValue string) (isValid bool) {
					total, _ := strconv.ParseFloat(attributeValue, 32)
					if total >= 0 {
						return true
					}
					return
				},
			},
			true,
		},
	}

	for _, tt := range tests {
		t.Logf("Case %s", tt.caseName)

		isValid, _ := tt.input.Validate()
		assert.Equal(t, isValid, tt.expected)
	}
}
