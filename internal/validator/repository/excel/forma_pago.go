package excel

import (
	"fmt"
	"github.com/xerardoo/invoice-cfdi-validator/internal/validator/entity"
	"github.com/xerardoo/invoice-cfdi-validator/pkg/datasource"
)

type FormaPagoExcelRepo struct {
	datasource.FileExcel
	sheetIndex    int
	sheetName     string
	columnIndex   int
	startRowIndex int
	endRowIndex   int
}

func NewFormaPagoExcelRepo(file datasource.FileExcel) FormaPagoExcelRepo {
	return FormaPagoExcelRepo{
		FileExcel: file,
	}
}

func (repo FormaPagoExcelRepo) SetSheet(sheetName string, sheetIndex int, columnIndex int, startRowIndex, endRowIndex int) FormaPagoExcelRepo {
	repo.sheetName = sheetName
	repo.sheetIndex = sheetIndex
	repo.columnIndex = columnIndex
	repo.startRowIndex = startRowIndex
	repo.endRowIndex = endRowIndex
	return repo
}

func (repo FormaPagoExcelRepo) GetAll() (formaPagos []entity.FormaPago, err error) {
	sheet, err := selectSheet(repo.File, repo.sheetIndex, repo.sheetName)
	if err != nil {
		return
	}

	for rowIndex := repo.startRowIndex; rowIndex <= repo.endRowIndex; rowIndex++ {
		cellValue, err := getCell(sheet, rowIndex, repo.columnIndex)
		if err != nil {
			return formaPagos, err
		}

		clave := fmt.Sprintf("%02v", cellValue)
		formaPagos = append(formaPagos, entity.FormaPago{clave})
	}
	return
}

func (repo FormaPagoExcelRepo) GetByClave(clave string) (formaPago entity.FormaPago, err error) {
	formaPagoList, err := repo.GetAll()
	if err != nil {
		return
	}

	for _, formaPago := range formaPagoList {
		if formaPago.Clave == clave {
			return formaPago, nil
		}
	}
	return
}
