package handler

import (
	"fmt"

	"g"
	"log"
	"models"

	"github.com/tealeg/xlsx"
)

func ImportExcel() []*models.CheckTable {

	excelFileName := g.Config().ImportExcelPath
	xlFile, err := xlsx.OpenFile(excelFileName)
	if err != nil {
		log.Fatal("read excel failed,err:%v", err.Error())
	}

	var excelTable = []*models.CheckTable{}
	for _, sheet := range xlFile.Sheets {
		for id, row := range sheet.Rows {
			if id == 0 {
				continue
			}
			cells := row.Cells
			var spanTime int
			spanTime, err = cells[6].Int()
			if err != nil {
				log.Fatal("bad span time,err:", err.Error(), cells[6].Value)
				spanTime = 8
			}

			threshold, err := cells[5].Float()
			if err != nil {
				log.Fatal("bad threshold value,err:", err.Error(), cells[5].Value)
				continue
			}
			var t = &models.CheckTable{
				ParScreen:   cells[0].Value,
				ChildScreen: cells[1].Value,
				Metric:      cells[2].Value,
				ComMode:     cells[3].Value,
				JudgeSymbol: cells[4].Value,
				Threshold:   threshold,
				SpanTime:    spanTime,
			}
			excelTable = append(excelTable, t)
		}
	}

	return excelTable
}

func ExportExcel(res []*models.CheckTable) {
	var file *xlsx.File
	var sheet *xlsx.Sheet
	var row *xlsx.Row
	var err error

	file = xlsx.NewFile()
	sheet, err = file.AddSheet("Sheet1")
	if err != nil {
		fmt.Printf(err.Error())
	}
	row = sheet.AddRow()
	cell0 := row.AddCell()
	cell0.Value = "巡检大类"

	cell1 := row.AddCell()
	cell1.Value = "子类"

	cell2 := row.AddCell()
	cell2.Value = "指标"

	cell3 := row.AddCell()
	cell3.Value = "统计策略"

	cell4 := row.AddCell()
	cell4.Value = "判断"

	cell5 := row.AddCell()
	cell5.Value = "告警阈值"

	cell6 := row.AddCell()
	cell6.Value = "时间段"

	cell7 := row.AddCell()
	cell7.Value = "是否异常"

	cell8 := row.AddCell()
	cell8.Value = "备注"

	for _, v := range res {
		row := sheet.AddRow()
		cell0 := row.AddCell()
		cell0.Value = v.ParScreen

		cell1 := row.AddCell()
		cell1.Value = v.ChildScreen

		cell2 := row.AddCell()
		cell2.Value = v.Metric

		cell3 := row.AddCell()
		cell3.Value = v.ComMode

		cell4 := row.AddCell()
		cell4.Value = v.JudgeSymbol

		cell5 := row.AddCell()
		cell5.Value = fmt.Sprintf("%f", v.Threshold)

		cell6 := row.AddCell()
		cell6.Value = fmt.Sprintf("%d", v.SpanTime)

		cell7 := row.AddCell()
		cell7.Value = fmt.Sprintf("%t", v.IsAbnormal)

		cell8 := row.AddCell()
		cell8.Value = v.Desc
	}

	err = file.Save(g.Config().ExportExecelPath)
	if err != nil {
		fmt.Println(err.Error())
	}
}
