package excel

import (
	"fmt"

	"github.com/guiphh/jira-helper/pkg/jclient"
	"github.com/tealeg/xlsx"
)

// WriteXlsx exports the issues to an excel file
func WriteXlsx(filename string, issues []jclient.MyIssue) {
	var file *xlsx.File
	var sheet *xlsx.Sheet
	var row *xlsx.Row
	var cell *xlsx.Cell

	file = xlsx.NewFile()
	sheet, err := file.AddSheet("Sheet1")
	if err != nil {
		fmt.Printf(err.Error())
	}

	for _, iss := range issues {
		row = sheet.AddRow()
		cell = row.AddCell()
		cell.Value = iss.Key
		cell = row.AddCell()
		cell.Value = iss.Summary
		cell = row.AddCell()
		cell.Value = iss.EpicKey
		cell = row.AddCell()
		cell.Value = iss.EpicName
	}

	err = file.Save(filename)
	if err != nil {
		fmt.Printf(err.Error())
	}
}
