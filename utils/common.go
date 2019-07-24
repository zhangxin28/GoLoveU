package utils

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"
	"strings"
	"sync"
	"time"
	"runtime/debug"
	"github.com/360EntSecGroup-Skylar/excelize"
)

func OpenExcel(file string) *excelize.File {
	f, err := excelize.OpenFile(file)
	if err != nil {
		CheckError(err)
	}
	return f
}

func GetSheetRowData(excelFile *excelize.File, sheetName string) ([][]string, int, int) {
	rows, _ := excelFile.GetRows(sheetName)
	rowCount := len(rows)
	columnCount := len(rows[0])
	return rows, rowCount, columnCount
}

func CheckError(err error) {
	if err != nil {
		fmt.Println(err)
		WaitUserEnterKeyToExit(false)
	}
}

func CreateExcel(folderName string, fileNameWithoutSuffix string, sheets ...string) {
	f := excelize.NewFile()
	for _, sheet := range sheets {
		f.NewSheet(sheet)
	}
	f.DeleteSheet("Sheet1") //delete the default sheet
	newFileIndex := time.Now().Format("20060102150405")
	err := f.SaveAs(fmt.Sprintf("./%s/%s_%s.xlsx", folderName, fileNameWithoutSuffix, newFileIndex))
	CheckError(err)
}

func WaitUserEnterKeyToExit(exit bool) {
	fmt.Println("\n\n按确认键退出...")
	reader := bufio.NewReader(os.Stdin)
	reader.ReadRune()
	//char,_,_ := reader.ReadRune()
	//fmt.Println(char)
	if exit == true {
		os.Exit(2)
	}
}

func GetFileName(file string) (fileName string, fileOnlyName string, fileSuffix string) {
	_, fileName = filepath.Split(file)
	fileSuffix = path.Ext(fileName)                         //获取文件后缀
	fileOnlyName = strings.TrimSuffix(fileName, fileSuffix) //获取文件名
	return fileName, fileOnlyName, fileSuffix
}

func GetFiles(filePrefix string) (files []string) {
	files = make([]string, 0, 5)
	currentDir, _ := filepath.Abs(filepath.Dir(os.Args[0]))

	dirList, err := ioutil.ReadDir(currentDir)
	CheckError(err)

	for _, v := range dirList {
		if !v.IsDir() {
			fileName := v.Name()
			if strings.HasPrefix(strings.ToUpper(fileName), filePrefix) {
				if strings.HasSuffix(strings.ToUpper(fileName), ".XLSX") ||
					strings.HasSuffix(strings.ToUpper(fileName), ".XLSM") {
					files = append(files, fileName)
				}
			}
		}
	}

	return files
}

func CreateNewFolder(path string) {
	if !CheckPathExists(path) {
		err := os.MkdirAll(fmt.Sprintf("./%s", path), os.ModePerm)
		CheckError(err)
	}
}

func CheckPathExists(path string) bool {
	_, err := os.Stat(path)
	return !os.IsNotExist(err)
}

func GetArrayIndex(value interface{}, compareFunc CompareFunc, values ...interface{}) int {
	result := -1
	for k, v := range values {
		if compareFunc == nil {
			if value == v {
				result = k
			}
		} else {
			if compareFunc(v, value) {
				result = k
			}
		}
	}
	return result
}

func GetMapKeys(m map[interface{}][]interface{}) (keys []interface{}) {
	keys = make([]interface{}, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	return keys
}

var mapGuard sync.Mutex

func GetSafeValue(f func() interface{}) (safeValue interface{}) {
	mapGuard.Lock()
	defer mapGuard.Unlock()
	safeValue = f()
	return safeValue
}

func PrintStack() {
	debug.PrintStack()
}
