package main

import (
	"fmt"
	"starbucks-tools/hexexport/hexchan"
	"starbucks-tools/hexexport/hexsource"
	"starbucks-tools/utils"
	"sync"
	"time"
)

func main() {
	toHandleFiles := utils.GetFiles(hexsource.HexSourceFileNamePrefix)
	if len(toHandleFiles) == 0 {
		fmt.Println("没有符合条件的HEX数据源文件")
		utils.WaitUserEnterKeyToExit(true)
	}

	start := time.Now()
	doHexCompute(toHandleFiles)
	fmt.Printf("\n计算Hex数据的消耗的时间为\t%s", time.Since(start))

	// resultFolderName := fmt.Sprintf("%s_%s", fileName, "RESULT")
	// utils.CreateNewFolder(resultFolderName)
	// fmt.Println()

	//utils.CreateExcel(resultFolderName, fileName, "test1", "test2", "test3")
	hexsource.DispatchHexDataSource()

	utils.WaitUserEnterKeyToExit(false)
}

func doHexCompute(files []string) {
	defer func() {
		if e := recover(); e != nil {
			utils.PrintStack()
		}
	}()

	var wg sync.WaitGroup
	dlChan := make(hexchan.DlChan)
	// step1 根据文件名找到EXCEL文件，然后生成需要计算的数据
	for _, file := range files {
		wg.Add(1)
		go func(f string) {
			//defer wg.Done()
			hexsource.GenerateHexData(f, dlChan)
		}(file)
	}

	// step2 对生成的数据进行计算
	go func() {
		for v := range dlChan {
			_, fileName, _ := utils.GetFileName(v.File)
			if v.DataSourceLength == 0 {
				fmt.Printf("文件:\t%s\t没有数据需要计算\n", fileName)
				fmt.Println()
				continue
			}

			fmt.Printf("正在计算文件\t%s\t中的数据", fileName)
			fmt.Printf("File %s has %d 条数据\n", fileName, v.DataSourceLength)
			wg.Done()
		}
	}()
	wg.Wait()
}
