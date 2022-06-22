package cfdi33

import (
	repository "github.com/xerardoo/invoice-cfdi-validator/internal/validator/repository/excel"
	"github.com/xerardoo/invoice-cfdi-validator/pkg/cfdi33"
	"github.com/xerardoo/invoice-cfdi-validator/pkg/datasource"
	"github.com/xerardoo/invoice-cfdi-validator/pkg/validator"
)

type FormaPagoUseCase struct {
	formaPagoRepo repository.FormaPagoExcelRepo
	cfdi          cfdi33.CFDI
}

func NewFormaPagoUseCase(fileExcel datasource.FileExcel, cfdi cfdi33.CFDI) *FormaPagoUseCase {
	formaPagoRepo := repository.NewFormaPagoExcelRepo(fileExcel)
	formaPagoRepo = formaPagoRepo.SetSheet("c_FormaPago", 0, 0, 6, 27)

	return &FormaPagoUseCase{
		formaPagoRepo: formaPagoRepo,
		cfdi:          cfdi,
	}
}

func (uc FormaPagoUseCase) FormaExists() validator.AttributeValidation {
	return validator.NewAttributeValidation(
		"FormaPago", uc.cfdi.FormaPago,
		"CFDI33104",
		"El campo FormaPago no contiene un valor del catálogo c_FormaPago",
		func(attributeValue string) bool {
			formaPago, err := uc.formaPagoRepo.GetByClave(attributeValue)
			if err != nil {
				return false
			}
			if formaPago.Clave != "" {
				return true
			}
			return false
		})
}
