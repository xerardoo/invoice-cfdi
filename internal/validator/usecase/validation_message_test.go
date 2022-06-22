package usecase_test

import (
	"github.com/stretchr/testify/assert"
	"github.com/xerardoo/invoice-cfdi-validator/internal/validator/usecase"
	"testing"
)

func TestCFDIValidationMessage_All(t *testing.T) {
	validationMessage := usecase.CFDIValidationMessage{}
	validationMessage.Cfdi.Version = "4.0"
	validationMessage.Cfdi.Validation.Code = "CDGI80"
	validationMessage.Cfdi.Validation.Result.IsValid = true
	validationMessage.Cfdi.Validation.Result.Message = "Valor no esta en catalogo."

	newValidationMsg := usecase.NewValidationMessage()
	newValidationMsg.SetVersion(validationMessage.Cfdi.Version)
	newValidationMsg.SetCode(validationMessage.Cfdi.Validation.Code)
	newValidationMsg.SetIsValid(validationMessage.Cfdi.Validation.Result.IsValid)
	newValidationMsg.SetMessage(validationMessage.Cfdi.Validation.Result.Message)

	assert.Equal(t, newValidationMsg.Cfdi.Version, validationMessage.Cfdi.Version)
	assert.Equal(t, newValidationMsg.Cfdi.Validation.Code, validationMessage.Cfdi.Validation.Code)
	assert.Equal(t, newValidationMsg.Cfdi.Validation.Result.IsValid, validationMessage.Cfdi.Validation.Result.IsValid)
	assert.Equal(t, newValidationMsg.Cfdi.Validation.Result.Message, validationMessage.Cfdi.Validation.Result.Message)
}
