package main

import (
	"fmt"
	"os"
	"starbucks-tools/hexexport/hexsource"
	"starbucks-tools/utils"
	"time"

	"github.com/cheggaaa/pb/v3"
)

func main() {
	toHandleFiles, err := utils.GetFiles(hexsource.HexSourceFileNamePrefix)
	utils.CheckError(err)
	if len(toHandleFiles) == 0 {
		fmt.Println("没有符合条件的HEX数据源文件")
		utils.WaitUserEnterKeyToExit(true)
	}

fileLoop:
	for _, file := range toHandleFiles {
		//hexData := hexsource.GetHexData(file)
		rowCount := hexsource.GenerateHexData("D:\\CodeSamples\\ZWGo\\src\\starbucks-tools\\hexexport\\ITEM_REQUEST_FORM_20190529_删除隐藏页签.xlsm")

		if rowCount == 0 {
			fmt.Printf("文件:\t%s\t没有数据需要计算\n", utils.GetFileName(file))
			fmt.Println()
			continue fileLoop
		}

		fmt.Println()
	}

	utils.WaitUserEnterKeyToExit(false)
}

func add(x, y int) {
	z := x + y
	fmt.Println(z)
}

func testPb() {
	count := 100000
	bar := pb.New(count)

	// refresh info every second (default 200ms)
	bar.SetRefreshRate(time.Second)

	// force set io.Writer, by default it's os.Stderr
	bar.SetWriter(os.Stdout)

	// start bar
	bar.Start()
}

func testProcess() {
	for i := 0; i != 10; i = i + 1 {
		fmt.Fprintf(os.Stdout, "result is %d\r", i)
		time.Sleep(time.Second * 1)
	}
	fmt.Println("over")
}
