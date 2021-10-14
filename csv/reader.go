package csv

import (
	"bytes"
	"errors"
	"fmt"
	"github.com/xuri/excelize/v2"
	"io"
	"os"
	"path/filepath"
)

type M map[string]string

func ExcelToMap(file *os.File, sheetName string) ([]M, error) {

	if filepath.Ext(file.Name()) != ".csv" && filepath.Ext(file.Name()) != ".xlsx" {
		return []M{}, errors.New("not excel or csv format")
	}

	buf := bytes.NewBuffer(nil)
	if _, err := io.Copy(buf, file); err != nil {
		return []M{}, err
	}

	xlsx, err := excelize.OpenReader(buf)
	if err != nil {
		return []M{}, err
	}

	rows := make([]M, 0)
	for i := 2; true; i++ {
		whitelistID, err := xlsx.GetCellValue(sheetName, fmt.Sprintf("A%d", i))
		if err != nil {
			return []M{}, err
		}

		NIK, err := xlsx.GetCellValue(sheetName, fmt.Sprintf("B%d", i))
		if err != nil {
			return []M{}, err
		}

		maxAmount, err := xlsx.GetCellValue(sheetName, fmt.Sprintf("C%d", i))
		if err != nil {
			return []M{}, err
		}

		maxUsage, err := xlsx.GetCellValue(sheetName, fmt.Sprintf("D%d", i))
		if err != nil {
			return []M{}, err
		}

		row := M{
			"whitelist_id": whitelistID,
			"nik":          NIK,
			"max_amount":   maxAmount,
			"max_usage":    maxUsage,
		}
		rows = append(rows, row)

		if cell, err := xlsx.GetCellValue(sheetName, fmt.Sprintf("A%d", i+1)); err != nil || cell == "" {
			break
		}
	}

	return rows, nil
}
