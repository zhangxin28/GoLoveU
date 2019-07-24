package main

import (
	"fmt"
	"starbucks-tools/hexexport/hexchan"
	"starbucks-tools/hexexport/hexsource"
	"starbucks-tools/utils"
)

func main() {
	toHandleFiles := utils.GetFiles(hexsource.HexSourceFileNamePrefix)
	if len(toHandleFiles) == 0 {
		fmt.Println("没有符合条件的HEX数据源文件")
		utils.WaitUserEnterKeyToExit(true)
	}
	doTest()
	//doHexCompute(toHandleFiles)

	// resultFolderName := fmt.Sprintf("%s_%s", fileName, "RESULT")
	// utils.CreateNewFolder(resultFolderName)
	// fmt.Println()

	//utils.CreateExcel(resultFolderName, fileName, "test1", "test2", "test3")
	hexsource.DispatchHexDataSource()

	utils.WaitUserEnterKeyToExit(false)
}

func doTest(){
	defer func() {
        if e := recover(); e != nil {
            utils.PrintStack()
        }
	}()

	zero := 0
    x := 3 / zero
    fmt.Println("x=", x)
}

func doHexCompute(files []string){
	defer func() {
        if e := recover(); e != nil {
            utils.PrintStack()
        }
	}()

	dlChan := make(hexchan.DlChan)
	// step1 根据文件名找到EXCEL文件，然后生成需要计算的数据
	for _, file := range files {
		go hexsource.GenerateHexData(file, dlChan)
	}
	close(dlChan)

	// step2 对生成的数据进行计算
	for v := range dlChan {
		_, fileName, _ := utils.GetFileName(v.File)
		if v.DataSourceLength == 0 {
			fmt.Printf("文件:\t%s\t没有数据需要计算\n", fileName)
			fmt.Println()
			continue
		}
	}
}
