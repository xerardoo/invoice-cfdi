package cfdi33

import (
	"fmt"
	repository "github.com/xerardoo/invoice-cfdi-validator/internal/validator/repository/excel"
	"github.com/xerardoo/invoice-cfdi-validator/pkg/cfdi33"
	"github.com/xerardoo/invoice-cfdi-validator/pkg/datasource"
	"github.com/xerardoo/invoice-cfdi-validator/pkg/validator"
)

type UsoCFDIUseCase struct {
	repo repository.UsoCFDIExcelRepo
	cfdi cfdi33.CFDI
}

func NewUsoCFDIUseCase(fileExcel datasource.FileExcel, cfdi cfdi33.CFDI) *UsoCFDIUseCase {
	repo := repository.NewUsoCFDIExcelRepo(fileExcel)
	repo = repo.SetSheet("c_UsoCFDI", 1, []int{0, 2, 3}, 6, 27)

	return &UsoCFDIUseCase{
		repo: repo,
		cfdi: cfdi,
	}
}

func (uc UsoCFDIUseCase) UsoCfdiReceptor() validator.AttributeValidation {
	return validator.NewAttributeValidation(
		"UsoCFDI", uc.cfdi.Receptor.UsoCFDI,
		"CFDI33141",
		"La clave del campo UsoCFDI debe corresponder con el tipo de persona (física o moral).",
		func(attributeValue string) bool {
			usoCfdi, err := uc.repo.GetByClave(attributeValue)
			if err != nil {
				fmt.Println("UsoCFDI err", err.Error())
				return false
			}
			switch len(uc.cfdi.Receptor.RFC) {
			case 12:
				if usoCfdi.IsMoral {
					return true
				}
				break
			case 13:
				if usoCfdi.IsFisica {
					return true
				}
				break
			}
			return false
		})
}
