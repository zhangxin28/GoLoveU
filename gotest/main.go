package main

import (
	"fmt"
	"starbucks-tools/gotest/structinterfacetest"
	"time"
)

func main() {
	testLongCoumpute()
}

//------------------------------------------------------------------------

func TestStructFuncTypeInterface() {
	structinterfacetest.HandleFunc("first default test", nil)
	structinterfacetest.HandleFunc("second test", func(message string) {
		fmt.Printf("This is the test with : %s\n", message)
	})
	structinterfacetest.HandleFunc("third test", func(message string) {
		fmt.Printf("This is the test with : %s\n", message)
	})
	structinterfacetest.DefaultRequestStruct.HandleFunc("fourth test", func(message string) {
		fmt.Printf("This is the test with : %s\n", message)
	})
	structinterfacetest.DefaultRequestStruct.HandleFunc("fifth test", structinterfacetest.FuncDoHandle(func(message string) {
		fmt.Printf("This is the test with : %s\n", message)
	}))
}

//------------------------------------------------------------------------

//------------------------------------------------------------------------
type simplePrint func(interface{})

var visitChan = make(chan string)

func visitCollection(sp simplePrint, items ...interface{}) {
	for _, item := range items {
		go sp(item)

		func(chanValue string) {
			fmt.Printf("Item ---[ %s ]---Visited", chanValue)
		}(<-visitChan)

	}

}
func testVisist() {
	visitCollection(simplePrint(func(a interface{}) {
		fmt.Printf("item type = %T\n", a)
		fmt.Println("item value = ", a)
		value := fmt.Sprintf("%s", a)
		visitChan <- value
	}), 1, 2, 3, 4, 5, []string{"one", "two", "three", "four"})
}

func testLongCoumpute() {
	a := 6.245614
	b := 2.718

	for index := 0; index < 10; index++ {
		testDuration(func() {
			for i := 0; i < 100000000; i++ {
				a = a + b
			}
		})
	}
}

func testDuration(f func()) {
	start := time.Now()

	f()

	fmt.Printf("函数执行时长为:\t%s\n", time.Since(start))
}

//------------------------------------------------------------------------
