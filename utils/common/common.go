package common

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"
	"runtime/debug"
	"strings"
	"sync"
)

// CompareFunc reprents a func type: (interface{}, interface{}) bool
type CompareFunc func(interface{}, interface{}) bool

// CheckError checks error
func CheckError(err error) {
	if err != nil {
		fmt.Println(err)
		WaitUserEnterKeyToExit(false)
	}
}

// WaitUserEnterKeyToExit waits the console applcation when use type enter
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

// GetFileName gets file name
func GetFileName(file string) (fileName string, fileOnlyName string, fileSuffix string) {
	_, fileName = filepath.Split(file)
	fileSuffix = path.Ext(fileName)                         //获取文件后缀
	fileOnlyName = strings.TrimSuffix(fileName, fileSuffix) //获取文件名
	return fileName, fileOnlyName, fileSuffix
}

// GetFiles gets files with the specific prefix
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

// CreateNewFolder creates a sub folder under the specific folder
func CreateNewFolder(path string) {
	if !CheckPathExists(path) {
		err := os.MkdirAll(fmt.Sprintf("./%s", path), os.ModePerm)
		CheckError(err)
	}
}

// CheckPathExists checks the path existed or not
func CheckPathExists(path string) bool {
	_, err := os.Stat(path)
	return !os.IsNotExist(err)
}

// GetArrayIndex returns the array index
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

// GetMapKeys returns the keys for a map
func GetMapKeys(m map[interface{}][]interface{}) (keys []interface{}) {
	keys = make([]interface{}, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	return keys
}

var mapGuard sync.Mutex

// GetSafeValue get value in some slice, under thread safe
func GetSafeValue(f func() interface{}) (safeValue interface{}) {
	mapGuard.Lock()
	defer mapGuard.Unlock()
	safeValue = f()
	return safeValue
}

// DoSafeSave do something, under thread safe
func DoSafeSave(f func()) {
	mapGuard.Lock()
	defer mapGuard.Unlock()
	f()
}

// PrintStack print the call stack
func PrintStack() {
	debug.PrintStack()
}
