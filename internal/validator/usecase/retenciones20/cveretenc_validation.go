package retenciones20

import (
	"github.com/xerardoo/invoice-cfdi-validator/pkg/retenciones20"
	"github.com/xerardoo/invoice-cfdi-validator/pkg/validator"
)

type CveRetencUseCase struct {
	retenciones retenciones20.Retenciones
}

func NewCveRetencUseCase(retenciones retenciones20.Retenciones) CveRetencUseCase {
	return CveRetencUseCase{retenciones: retenciones}
}

func (cr CveRetencUseCase) CveRetencExists() validator.AttributeValidation {
	return validator.NewAttributeValidation(
		"CveRetenc", cr.retenciones.CveRetenc,
		"Reten20106",
		"El campo DescRetenc debe existir.",
		func(attributeValue string) bool {
			if attributeValue == "25" && cr.retenciones.DescRetenc == "" {
				return false
			}
			return true
		})
}
