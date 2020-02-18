package services

import (
	"fmt"
	"github.com/360EntSecGroup-Skylar/excelize/v2"
)

func GetExcelData(file string) (rows [][]string, err error) {
	fmt.Println(file)
	xlsx, err := excelize.OpenFile(file)
	if err != nil {
		fmt.Println("read excel:", err)
	}
	rows, err = xlsx.GetRows("Sheet1")
	return
}
