package hexsource

import (
	"fmt"
	"starbucks-tools/utils"
	"time"

	"github.com/360EntSecGroup-Skylar/excelize"
)

var HexSourceFileNamePrefix = "ITEM_REQUEST_FORM"

var AllHexFields = []string{"商品编码", "SBI", "商品名称", "商品英文名称", "核算单位", "税别", "是否增值税商品", "是否可主配", "商品类别编码", "存储类型", "行类型", "产品属性", "Oracle凭证科目", "BOM类型", "适用开始日", "条码二", "保质期（天）", "是否可调拨", "是否可调整", "是否可盘点", "供应商编码", "采购价格(含税)", "应计税率", "是否主供应商", "是否可采购", "是否可销售", "销售价格", "是否可订货", "计划到货天数", "最小订量", "最大订量", "是否可订货", "计划到货天数", "最小订量", "最大定量"}
var resultAreaFields = map[string][]string{
	"商品基本资料":   []string{"商品编码", "SBI", "商品名称", "商品英文名称", "核算单位", "税别", "是否增值税商品", "是否可主配", "商品类别编码", "存储类型", "行类型", "产品属性", "用友凭证科目", "BOM类型", "适用开始日", "条码二", "保质期（天）"},
	"区域商品属性":   []string{"属性区域编码", "商品编码", "生效开始日期", "生效结束日期", "是否可调拨", "是否可调整", "是否可盘点"},
	"区域商品采购价格": []string{"采购区域编码", "供应商编码", "商品编码", "生效开始日期", "生效结束日期", "采购价格(含税)", "应计税率", "是否主供应商", "是否可采购"},
	"区域商品售价":   []string{"市场区域编码", "商品编码", "生效开始日期", "生效结束日期", "销售价格", "是否可销售"},
	"区域商品配送规则": []string{"配送区域编码", "商品编码", "生效开始日期", "生效结束日期", "物流模式", "配送中心", "是否可订货", "周一", "周二", "周三", "周四", "周五", "周六", "周日", "计划到货天数", "最小订量", "最大订量", "递增订量"},
	"区域商品配方":   []string{"BOM区域", "BOM商品编码", "生效开始时间", "生效结束时间", "原料商品编码", "配方数量"},
}

var hexDataSource = make(map[string][]string, len(AllHexFields))

func GenerateHexData(file string) int {
	fileName := utils.GetFileName(file)
	start := time.Now()
	f, err := excelize.OpenFile(file)
	if err != nil {
		utils.CheckError(err)
	}

	fmt.Printf("正在打开文件:%s,\t耗时:%s\n", fileName, time.Since(start))
	start = time.Now()
	// Get all the rows in the Sheet1.

	rows, _ := f.GetRows("HEX")
	rowCount := len(rows)
	columnCount := len(rows[0])
	//第一条显示为：商品基本资料，不做考虑
	//第二条显示为：列的表头，不做考虑
	if rowCount-2 <= 0 {
		return 0
	}

	for columnIndex := 0; columnIndex < columnCount; columnIndex++ {
		columnHeadValue := rows[1][columnIndex]
		columnValues := make([]string, rowCount-2)
		for rowIndex := 2; rowIndex < rowCount; rowIndex++ {
			columnValues[rowIndex-2] = rows[rowIndex][columnIndex]
		}
		hexDataSource[columnHeadValue] = columnValues
	}
	mapOneRowLength := len(hexDataSource[AllHexFields[0]])
	fmt.Printf("正在生成数据,耗时:%s, \t数据总计:%d条\n", time.Since(start), mapOneRowLength)

	return mapOneRowLength
}

func DispatchHexDataSource(areaName string) (string, [][]string) {
	return "", nil
}

func testExcelCreate() {
	f := excelize.NewFile()
	// Create a new sheet.
	index := f.NewSheet("Sheet2")
	// Set value of a cell.
	f.SetCellValue("Sheet2", "A2", "Hello world.")
	f.SetCellValue("Sheet1", "B2", 100)
	// Set active sheet of the workbook.
	f.SetActiveSheet(index)
	// Save xlsx file by the given path.
	err := f.SaveAs("./Book1.xlsx")
	if err != nil {
		fmt.Println(err)
	}
}
