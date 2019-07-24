package hexsource

import (
	"fmt"
	"starbucks-tools/hexexport/hexchan"
	"starbucks-tools/utils"
	"strings"
	"sync"
	"time"
)

var HexSourceFileNamePrefix = "ITEM_REQUEST_FORM"

var allHexHeaderFields = map[int]string{
	0:  "商品编码",
	1:  "SBI",
	2:  "商品名称",
	3:  "商品英文名称",
	4:  "核算单位",
	5:  "税别",
	6:  "是否增值税商品",
	7:  "是否可主配",
	8:  "商品类别编码",
	9:  "存储类型",
	10: "行类型",
	11: "产品属性",
	12: "Oracle凭证科目",
	13: "BOM类型",
	14: "适用开始日",
	15: "条码二",
	16: "保质期（天）",
	17: "是否可调拨",
	18: "是否可调整",
	19: "是否可盘点",
	20: "供应商编码",
	21: "采购价格(含税)",
	22: "应计税率",
	23: "是否主供应商",
	24: "是否可采购",
	25: "是否可销售",
	26: "销售价格",
	27: "是否可订货-F04",
	28: "计划到货天数-F04",
	29: "最小订量-F04",
	30: "最大订量-F04",
	31: "是否可订货-F05",
	32: "计划到货天数-F05",
	33: "最小订量-F05",
	34: "最大定量-F05",
}

func GenerateHexData(file string, dlChan hexchan.DlChan) {
	fileName, _, _ := utils.GetFileName(file)
	start := time.Now()
	f := utils.OpenExcel(file)

	fmt.Printf("文件【%s】正在打开,\t耗时:%s\n", fileName, time.Since(start))
	start = time.Now()

	//最终生成的该文件中含有多少条数据
	dataLength := 0

	rows, rowCount, columnCount := utils.GetSheetRowData(f, "HEX")
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
				headerValues[columnHeadValue] = columnValues
			}
		}

		dataLength = len(headerValues[allHexHeaderFields[0]])
		fmt.Printf("文件【%s】正在组装数据,\t数据总计:%d条,\t耗时:%s\n", fileName, dataLength, time.Since(start))
	}

	dlChan <- hexchan.DlChanStruct{File: file, DataSourceLength: dataLength, HexData: headerValues}
}

func DispatchHexDataSource(dls hexchan.DlChanStruct) {
	if dls.DataSourceLength == 0 {
		fmt.Printf("文件【%s】没有数据需要计算\n", dls.File)
		fmt.Println()
	}

	fmt.Printf("文件【%s】的数据正在计算\n", dls.File)
	fmt.Println()
}

var hexDsGuard sync.Mutex
