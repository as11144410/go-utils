package helper

import (
	"bufio"
	"encoding/csv"
	"golang.org/x/text/encoding/simplifiedchinese"
	"golang.org/x/text/transform"
	"log"
	"os"
)

// CommonExportCsv 导出csv
func CommonExportCsv(fileName string, data [][]string) error {
	fp, err := os.Create(fileName) // 创建文件句柄
	if err != nil {
		log.Fatalf("创建文件["+fileName+"]句柄失败,%v", err)
		return err
	}
	defer fp.Close()
	fp.WriteString("\xEF\xBB\xBF") // 写入UTF-8 BOM
	w := csv.NewWriter(fp)         //创建一个新的写入文件流
	w.WriteAll(data)
	w.Flush()
	return nil
}

// LoadCsvFileContent 读取csv数据返回
func LoadCsvFileContent(dir string) [][]string {
	rs := make([][]string, 0)
	if dir == "" {
		return rs
	}
	f, err := os.Open(dir)
	if err != nil {
		return rs
	}
	defer f.Close()
	reader1 := transform.NewReader(bufio.NewReader(f), simplifiedchinese.GBK.NewDecoder())
	reader := csv.NewReader(reader1)
	rs, err = reader.ReadAll()
	return rs
}
