package cfdi33

import (
	"github.com/xerardoo/invoice-cfdi-validator/pkg/cfdi33"
	"github.com/xerardoo/invoice-cfdi-validator/pkg/validator"
	"regexp"
)

type FechaUseCase struct {
	Cfdi cfdi33.CFDI
}

func NewFechaUseCase(cfdi cfdi33.CFDI) *FechaUseCase {
	return &FechaUseCase{Cfdi: cfdi}
}

func (f FechaUseCase) FechaFormat() validator.AttributeValidation {
	return validator.NewAttributeValidation(
		"Fecha", f.Cfdi.Fecha,
		"CFDI33101",
		"El campo Fecha no cumple con el patrón requerido.",
		func(attributeValue string) bool {
			fechaPattern := `^(20[1-9][0-9])-(0[1-9]|1[0-2])-(0[1-9]|[12][0-9]|3[01])T(([01][0-9]|2[0-3]):[0-5][0-9]:[0-5][0-9])$`
			matched, _ := regexp.MatchString(fechaPattern, attributeValue)
			return matched
		})
}
