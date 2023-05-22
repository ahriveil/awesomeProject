package main

import (
	"github.com/xuri/excelize/v2"
	"log"
)

func main() {
	headLine := []string{"岗位名称"}
	excel, err := WriteFile(headLine)
	if err != nil {
		log.Printf("error:%s\n", err)
		return
	}
	err = excel.Save("职位表格")
	if err != nil {
		log.Printf("error:%s\n", err)
		return
	}

}

func writeFileExcel() {
	//保存数据
	f := excelize.NewFile()
	// 创建一个工作表
	index := f.NewSheet("b站职位")
	// 设置单元格的值
	f.SetCellValue("Sheet2", "A2", "Hello world.")
	f.SetCellValue("Sheet1", "B2", 100)
	// 设置工作簿的默认工作表
	f.SetActiveSheet(index)
	// 根据指定路径保存文件
	if err := f.SaveAs("/Users/tianyu06/test/Book1.xlsx"); err != nil {
		println(err.Error())
	}

}
