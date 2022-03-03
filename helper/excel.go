package helper

import (
	"fmt"
	"github.com/360EntSecGroup-Skylar/excelize"
	"strconv"
)

// CommonExport 导出excel
func CommonExport(title []string, fileName string, data [][]interface{}) error {
	xlsx := excelize.NewFile()
	// 设置表头
	var i = 'A'
	for _, val := range title {
		c := fmt.Sprintf("%c", i)
		xlsx.SetCellValue("Sheet1", c+"1", val)
		i++
	}
	// 填充数据
	var k = 2
	for _, va := range data {
		var j = 'A'
		for _, v := range va {
			d := fmt.Sprintf("%c", j) + strconv.Itoa(k)
			xlsx.SetCellValue("Sheet1", d, v)
			j++
		}
		k++
	}
	err := xlsx.SaveAs(fileName)
	return err
}

// LoadExcelFileContent 读取excel数据返回
func LoadExcelFileContent(dir string, sheet string) [][]string {
	rs := make([][]string, 0)
	if dir == "" {
		return rs
	}
	f, err := excelize.OpenFile(dir)
	if err != nil {
		return rs
	}
	rs = f.GetRows(sheet)
	return rs
}
