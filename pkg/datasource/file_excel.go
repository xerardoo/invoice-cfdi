package datasource

import (
	"github.com/shakinm/xlsReader/xls"
)

type FileExcel struct {
	Filename string
	File     xls.Workbook
}

func NewFileExcel(filename string) FileExcel {
	return FileExcel{Filename: filename}
}

func (f *FileExcel) Connect() (err error) {
	f.File, err = xls.OpenFile(f.Filename)
	if err != nil {
		return
	}
	return
}
