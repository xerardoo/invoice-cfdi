package retenciones20

import (
	"github.com/xerardoo/invoice-cfdi-validator/pkg/retenciones20"
	"github.com/xerardoo/invoice-cfdi-validator/pkg/validator"
	"strconv"
)

type MesIniUseCase struct {
	Retenciones retenciones20.Retenciones
}

func NewMesIniUseCase(retenciones retenciones20.Retenciones) *MesIniUseCase {
	return &MesIniUseCase{Retenciones: retenciones}
}

func (mi MesIniUseCase) MesIniValidation() validator.AttributeValidation {
	return validator.NewAttributeValidation(
		"Retenciones:Periodo:MesIni", mi.Retenciones.Periodo.MesIni,
		"Reten20122",
		"El campo MesIni no es menor o igual que el campo MesFin.",
		func(attributeValue string) (isValid bool) {
			mesIni, _ := strconv.Atoi(attributeValue)
			mesFin, _ := strconv.Atoi(mi.Retenciones.Periodo.MesFin)

			if mesFin >= mesIni {
				return true
			}
			return
		})
}
