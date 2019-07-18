package main

import (
	"fmt"
	"starbucks-tools/hexexport/hexsource"
	"starbucks-tools/utils"
)

func main() {
	toHandleFiles := utils.GetFiles(hexsource.HexSourceFileNamePrefix)
	if len(toHandleFiles) == 0 {
		fmt.Println("没有符合条件的HEX数据源文件")
		utils.WaitUserEnterKeyToExit(true)
	}

fileLoop:
	for _, file := range toHandleFiles {
		//hexData := hexsource.GetHexData(file)
		rowCount := hexsource.GenerateHexData(file)

		_, fileName, _ := utils.GetFileName(file)
		if rowCount == 0 {
			fmt.Printf("文件:\t%s\t没有数据需要计算\n", fileName)
			fmt.Println()
			continue fileLoop
		}

		resultFolderName := fmt.Sprintf("%s_%s", fileName, "RESULT")
		utils.CreateNewFolder(resultFolderName)
		fmt.Println()
	}

	//utils.CreateExcel(resultFolderName, fileName, "test1", "test2", "test3")
	hexsource.DispatchHexDataSource()

	utils.WaitUserEnterKeyToExit(false)
}
