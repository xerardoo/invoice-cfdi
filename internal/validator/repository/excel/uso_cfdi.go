package excel

import (
	"github.com/xerardoo/invoice-cfdi-validator/internal/validator/entity"
	"github.com/xerardoo/invoice-cfdi-validator/pkg/datasource"
	"strings"
)

type UsoCFDIExcelRepo struct {
	datasource.FileExcel
	sheetIndex    int
	sheetName     string
	columnIndexes []int
	startRowIndex int
	endRowIndex   int
}

func NewUsoCFDIExcelRepo(file datasource.FileExcel) UsoCFDIExcelRepo {
	return UsoCFDIExcelRepo{
		FileExcel: file,
	}
}

func (repo UsoCFDIExcelRepo) SetSheet(sheetName string, sheetIndex int, columnIndexes []int, startRowIndex, endRowIndex int) UsoCFDIExcelRepo {
	repo.sheetName = sheetName
	repo.sheetIndex = sheetIndex
	repo.columnIndexes = columnIndexes
	repo.startRowIndex = startRowIndex
	repo.endRowIndex = endRowIndex
	return repo
}

func (repo UsoCFDIExcelRepo) GetAll() (usos []entity.UsoCFDI, err error) {
	sheet, err := selectSheet(repo.File, repo.sheetIndex, repo.sheetName)
	if err != nil {
		return
	}

	for rowIndex := repo.startRowIndex; rowIndex <= repo.endRowIndex; rowIndex++ {
		currentRow := []string{}

		for _, columIndex := range repo.columnIndexes {
			cellValue, err := getCell(sheet, rowIndex, columIndex)
			if err != nil {
				return usos, err
			}

			currentRow = append(currentRow, cellValue)
		} // foreach

		uso := parseToUsoCFDI(currentRow)
		usos = append(usos, uso)
	} // for
	return
}

func (repo UsoCFDIExcelRepo) GetByClave(clave string) (uso entity.UsoCFDI, err error) {
	usosList, err := repo.GetAll()
	if err != nil {
		return
	}

	for _, uso := range usosList {
		if uso.Clave == clave {
			return uso, nil
		}
	}
	return
}

func parseToUsoCFDI(row []string) (uso entity.UsoCFDI) {
	clave := row[0]
	fisica := row[1]
	moral := row[2]

	uso = entity.UsoCFDI{
		Clave:    clave,
		IsFisica: !strings.Contains(fisica, "No"),
		IsMoral:  !strings.Contains(moral, "No"),
	}
	return
}
