package main

import (
	"fmt"
	"starbucks-tools/hexexport/hexchan"
	"starbucks-tools/hexexport/hexsource"
	"starbucks-tools/utils"
	//"sync"
)

func main() {
	// 获取符合条件的HEX数据源文件
	// 文件名以ITEM_REQUEST_FORM打头
	toHandleFiles := utils.GetFiles(hexsource.HexSourceFileNamePrefix)
	if len(toHandleFiles) == 0 {
		fmt.Println("没有符合条件的HEX数据源文件")
		utils.WaitUserEnterKeyToExit(true)
	}	

	// 并行执行流程：对多文件生成HEX数据
	// 并将数据保存在通道中
	hexGenerateChan := make(hexchan.HexGenerateChan)
	for _, file := range toHandleFiles {
		go func(file string) {
			hexsource.GenerateHexData(file, hexGenerateChan)
		}(file)
	}

	// 通过通道拿到生成的HEX数据，并将该数据保存
	hexGenerateDataMap := make(map[string]map[string][]string)
	jobsLeft := len(toHandleFiles)
	for hgr := range hexGenerateChan {
		jobsLeft--
		utils.DoSafeSave(func() {
			hexGenerateDataMap[hgr.File] = hgr.HexData
		})
		if jobsLeft == 0 {
			break
		}
	}
	close(hexGenerateChan)

	// 通过通道拿到计算出来的HEX数据	
	// 对分配的数据进行并行计算
	// 并将数据保存在通道中
	computeResultChan := make(chan int)	
	for file := range hexGenerateDataMap {
		go func(file string){
			hexsource.DispatchHexDataSource(file, hexGenerateDataMap[file], computeResultChan)
		}(file)		
	}

	// 获取通道中的HexCompute的结果数据
	jobsLeft = len(hexGenerateDataMap)
	// 清空之前保存的HEX待计算数据
	hexGenerateDataMap = make(map[string]map[string][]string)
	for v := range computeResultChan {
		jobsLeft--
		fmt.Println("Hex Compute", v)
		if jobsLeft == 0 {
			break
		}
	}
	close(computeResultChan)

	utils.WaitUserEnterKeyToExit(false)
}

// func main() {
// 	toHandleFiles := utils.GetFiles(hexsource.HexSourceFileNamePrefix)
// 	if len(toHandleFiles) == 0 {
// 		fmt.Println("没有符合条件的HEX数据源文件")
// 		utils.WaitUserEnterKeyToExit(true)
// 	}

// 	start := time.Now()
// 	var wg sync.WaitGroup
// 	dlChan := make(hexchan.DlChan)
// 	// step1 根据文件名找到EXCEL文件，然后生成需要计算的数据
// 	for _, file := range toHandleFiles {
// 		wg.Add(1)
// 		go func(f string) {
// 			//defer wg.Done()
// 			hexsource.GenerateHexData(f, dlChan)
// 		}(file)
// 	}

// 	// step2 对生成的数据进行计算
// 	go func() {
// 		for v := range dlChan {
// 			go func(dls hexchan.DlChanStruct) {
// 				defer wg.Done()
// 				hexsource.DispatchHexDataSource(dls)
// 			}(v)
// 		}
// 	}()
// 	wg.Wait()

// 	fmt.Printf("\n计算Hex数据的消耗的时间为\t%s", time.Since(start))

// 	// resultFolderName := fmt.Sprintf("%s_%s", fileName, "RESULT")
// 	// utils.CreateNewFolder(resultFolderName)
// 	// fmt.Println()

// 	//utils.CreateExcel(resultFolderName, fileName, "test1", "test2", "test3")
// 	//hexsource.DispatchHexDataSource(toHandleFiles[0])

// 	// for v := range hexResultChan{
// 	// 	fmt.Println("HEXRESULT ",v)
// 	// }

// 	utils.WaitUserEnterKeyToExit(false)
// }

// func doHexCompute(files []string) {
// 	var wg sync.WaitGroup
// 	dlChan := make(hexchan.HexGenerateChan)
// 	// step1 根据文件名找到EXCEL文件，然后生成需要计算的数据
// 	for _, file := range files {
// 		wg.Add(1)
// 		go func(f string) {
// 			//defer wg.Done()
// 			hexsource.GenerateHexData(f, dlChan)
// 		}(file)
// 	}

// 	c := make(chan int)
// 	// step2 对生成的数据进行计算
// 	go func() {
// 		for v := range dlChan {
// 			go func(dls hexchan.HexGenerateResult) {
// 				defer wg.Done()
// 				hexsource.DispatchHexDataSource(dls, c)
// 			}(v)
// 		}
// 	}()
// 	wg.Wait()
// }
