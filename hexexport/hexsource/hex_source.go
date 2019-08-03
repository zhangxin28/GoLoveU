package hexsource

import (
	"fmt"
	"starbucks/tools/hextool/excelhandler"
	"starbucks/tools/hextool/hexchan"
	"starbucks/tools/utils/common"
	"strings"
	"time"
)

// HexSourceFileNamePrefix reprents the hex soure file prefix anme
var HexSourceFileNamePrefix = "ITEM_REQUEST_FORM"

// HeaderSourceResultMap represents the hex source header mapping

type HeaderSourceResultMap struct {
	SourceName string
	ResultName string
}

var allHexHeaderFields = map[int]HeaderSourceResultMap{
	0:  HeaderSourceResultMap{"商品编码", ""},
	1:  HeaderSourceResultMap{"SBI", ""},
	2:  HeaderSourceResultMap{"商品名称", ""},
	3:  HeaderSourceResultMap{"商品英文名称", ""},
	4:  HeaderSourceResultMap{"核算单位", ""},
	5:  HeaderSourceResultMap{"税别", ""},
	6:  HeaderSourceResultMap{"是否增值税商品", ""},
	7:  HeaderSourceResultMap{"是否可主配", ""},
	8:  HeaderSourceResultMap{"商品类别编码", ""},
	9:  HeaderSourceResultMap{"存储类型", ""},
	10: HeaderSourceResultMap{"行类型", ""},
	11: HeaderSourceResultMap{"产品属性", ""},
	12: HeaderSourceResultMap{"Oracle凭证科目", "用友凭证科目"},
	13: HeaderSourceResultMap{"BOM类型", ""},
	14: HeaderSourceResultMap{"适用开始日", ""},
	15: HeaderSourceResultMap{"条码二", ""},
	16: HeaderSourceResultMap{"保质期（天）", ""},
	17: HeaderSourceResultMap{"是否可调拨", ""},
	18: HeaderSourceResultMap{"是否可调整", ""},
	19: HeaderSourceResultMap{"是否可盘点", ""},
	20: HeaderSourceResultMap{"供应商编码", ""},
	21: HeaderSourceResultMap{"采购价格(含税)", ""},
	22: HeaderSourceResultMap{"应计税率", ""},
	23: HeaderSourceResultMap{"是否主供应商", ""},
	24: HeaderSourceResultMap{"是否可采购", ""},
	25: HeaderSourceResultMap{"是否可销售", ""},
	26: HeaderSourceResultMap{"销售价格", ""},
	27: HeaderSourceResultMap{"是否可订货-F04", "是否可订货"},
	28: HeaderSourceResultMap{"计划到货天数-F04", "计划到货天数"},
	29: HeaderSourceResultMap{"最小订量-F04", "最小订量"},
	30: HeaderSourceResultMap{"最大订量-F04", "最大订量"},
	31: HeaderSourceResultMap{"是否可订货-F05", "是否可订货"},
	32: HeaderSourceResultMap{"计划到货天数-F05", "计划到货天数"},
	33: HeaderSourceResultMap{"最小订量-F05", "最小订量"},
	34: HeaderSourceResultMap{"最大定量-F05", "最大定量"},
}

// GenerateHexData reprents the logic to get the hex source data
func GenerateHexData(file string, dlChan hexchan.HexGenerateChan) {
	fileName, _, _ := common.GetFileName(file)
	start := time.Now()
	f, err := excelhandler.OpenExcel(file)
	common.CheckError(err)

	fmt.Printf("文件【%s】正在打开,\t耗时:%s\n", fileName, time.Since(start))
	start = time.Now()

	rows, rowCount, columnCount := excelhandler.GetSheetRowData(f, "HEX")
	headerValues := make(map[string][]string, columnCount)

	//第一条显示为：商品基本资料，不做考虑
	//第二条显示为：列的表头，不做考虑
	if rowCount-2 > 0 {
		for columnIndex := 0; columnIndex < columnCount; columnIndex++ {
			if columnHeadValue, ok := allHexHeaderFields[columnIndex]; ok {
				columnValues := make([]string, 0)
				for rowIndex := 2; rowIndex < rowCount; rowIndex++ {
					if firstColumnValue := strings.Trim(rows[rowIndex][0], " "); firstColumnValue != "" {
						columnValues = append(columnValues, strings.Trim(rows[rowIndex][columnIndex], " "))
					}
				}
				headerValues[columnHeadValue.SourceName] = columnValues
			}
		}

		fmt.Printf("文件【%s】正在组装数据,\t数据总计:%d条,\t耗时:%s\n", fileName, len(headerValues[allHexHeaderFields[0].SourceName]), time.Since(start))
	}

	dlChan <- hexchan.HexGenerateResult{File: file, HexData: headerValues}
}

// DispatchHexDataSource reprents the logic to compute and dispatch the hex result data
func DispatchHexDataSource(file string, hexData map[string][]string, c chan int) {
	if len(hexData[allHexHeaderFields[0].SourceName]) == 0 {
		fmt.Printf("文件【%s】没有数据需要计算\n", file)
		fmt.Println()
	}

	fmt.Printf("文件【%s】的数据正在计算，拥有数据%d条\n", file, len(hexData[allHexHeaderFields[0].SourceName]))
	fmt.Println()
	c <- 1
}
