package excelhandler

import (
	"fmt"
	"time"

	"github.com/360EntSecGroup-Skylar/excelize"
)

// OpenExcel opens the excel file
func OpenExcel(file string) (*excelize.File, error) {
	f, err := excelize.OpenFile(file)
	return f, err
}

// GetSheetRowData returns sheet rows
func GetSheetRowData(excelFile *excelize.File, sheetName string) (rows [][]string, rowCount int, columnCount int) {
	rows = excelFile.GetRows(sheetName)
	rowCount = len(rows)
	columnCount = len(rows[0])
	return rows, rowCount, columnCount
}

// CreateExcel creates the excel file
func CreateExcel(folderName string, fileNameWithoutSuffix string, sheets ...string) error {
	f := excelize.NewFile()
	for _, sheet := range sheets {
		f.NewSheet(sheet)
	}
	f.DeleteSheet("Sheet1") //delete the default sheet
	newFileIndex := time.Now().Format("20060102150405")
	err := f.SaveAs(fmt.Sprintf("./%s/%s_%s.xlsx", folderName, fileNameWithoutSuffix, newFileIndex))
	return err
}
