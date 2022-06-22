package retenciones20

import (
	"github.com/xerardoo/invoice-cfdi-validator/pkg/retenciones20"
	"github.com/xerardoo/invoice-cfdi-validator/pkg/validator"
	"regexp"
)

type FechaExpUseCase struct {
	retenciones retenciones20.Retenciones
}

func NewFechaExpUseCase(retenciones retenciones20.Retenciones) FechaExpUseCase {
	return FechaExpUseCase{retenciones}
}

func (fe FechaExpUseCase) FechaExpFormat() validator.AttributeValidation {
	return validator.NewAttributeValidation(
		"FechaExp", fe.retenciones.FechaExp,
		"Reten20103",
		"El campo FechaExp no cumple con el patrón requerido.",
		func(attributeValue string) bool {
			fechaPattern := `^(20[1-9][0-9])-(0[1-9]|1[0-2])-(0[1-9]|[12][0-9]|3[01])T(([01][0-9]|2[0-3]):[0-5][0-9]:[0-5][0-9])$`
			matched, _ := regexp.MatchString(fechaPattern, attributeValue)
			return matched
		})
}
