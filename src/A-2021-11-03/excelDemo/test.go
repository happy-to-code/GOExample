package main

import (
	"fmt"

	"github.com/360EntSecGroup-Skylar/excelize"
)

func main() {
	f, err := excelize.OpenFile("E:\\20.06.16Project\\GOExample\\src\\A-2021-11-03\\excelDemo\\file\\err.xlsx")
	if err != nil {
		fmt.Println(err)
		return
	}
	// Get value from cell by given worksheet name and axis.
	// cell := f.GetCellValue("errlog_rebuild", "B2")
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }
	// fmt.Println(cell)
	rows := f.GetRows("errlog_rebuild")
	for _, row := range rows {
		for _, colCell := range row {
			fmt.Print(colCell, "||||")
		}
		fmt.Println()
	}
}
