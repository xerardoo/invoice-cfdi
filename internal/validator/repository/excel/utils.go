package excel

import (
	"errors"
	"github.com/shakinm/xlsReader/xls"
)

var (
	errSheetNotFound = errors.New("sheet not found")
)

func selectSheet(workbook xls.Workbook, sheetIndex int, sheetName string) (sheet *xls.Sheet, err error) {
	sheet, err = workbook.GetSheet(sheetIndex)
	if err != nil {
		return
	}

	currentSheetName := sheet.GetName()
	if currentSheetName != sheetName {
		err = errSheetNotFound
		return
	}
	return
}

func getCell(sheet *xls.Sheet, row, column int) (value string, err error) {
	currentRow, err := sheet.GetRow(row)
	if err != nil {
		return
	}
	currentCol, err := currentRow.GetCol(column)
	if err != nil {
		return
	}
	value = currentCol.GetString()
	return
}
