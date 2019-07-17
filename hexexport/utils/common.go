package utils

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func CheckError(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
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

func GetFileName(file string) string {
	_, fileName := filepath.Split(file)
	return fileName
}

func GetFiles(filePrefix string) (files []string, err error) {
	files = make([]string, 0, 5)
	currentDir, _ := filepath.Abs(filepath.Dir(os.Args[0]))

	err = filepath.Walk(currentDir, func(filename string, fi os.FileInfo, err error) error { //遍历目录
		//if err != nil { //忽略错误
		// return err
		//}

		if fi.IsDir() { // 忽略目录
			return nil
		}

		if strings.HasPrefix(strings.ToUpper(fi.Name()), filePrefix) {
			if strings.HasSuffix(strings.ToUpper(fi.Name()), ".XLSX") ||
				strings.HasSuffix(strings.ToUpper(fi.Name()), ".XLSM") {
				files = append(files, filename)
			}
		}

		return nil
	})

	return files, err
}

func CheckPathExists(path string) bool {
	_, err := os.Stat(path)
	return !os.IsNotExist(err)
}

func GetArrayIndex(value interface{}, compareFunc CompareFunc, values ...interface{}) int {
	result := 0
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
