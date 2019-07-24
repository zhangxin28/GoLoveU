package main

import (
	"fmt"	
    "io/ioutil"
    "log"
    "net/http"
	"sync"
	"time"
	"starbucks-tools/gotest/structinterfacetest"
	"starbucks-tools/utils"
)

func main() {
	testChan()


	utils.WaitUserEnterKeyToExit(false)
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

//------------------------------------------------------------------------
type responseStruct struct {
	url string
	jsonResponses string
	duration string
}
func testChan() {
	urls := []string{
        "http://api.douban.com/v2/book/isbn/9787218087351",
        "http://ip.taobao.com/service/getIpInfo.php?ip=202.101.172.35",
        "https://jsonplaceholder.typicode.com/todos/1",
    }
    jsonResponses := make(chan responseStruct)
    var wg sync.WaitGroup
    wg.Add(len(urls))
    for _, url := range urls {
        go func(url string) {
			defer wg.Done()
			start := time.Now()
            res, err := http.Get(url)
            if err != nil {
                log.Fatal(err)
            } else {
                defer res.Body.Close()
                body, err := ioutil.ReadAll(res.Body)
                if err != nil {
                    log.Fatal(err)
                } else {
                    jsonResponses <- responseStruct {url, string(body), fmt.Sprintf("%s",time.Since(start)) }
                }
            }
        }(url)
    }
    go func() {
        for response := range jsonResponses {
            fmt.Printf("Get Response From Url %s\nReponse Body %s\nReponse Duration %s\n\n",response.url,response.jsonResponses,response.duration)
		}		
    }()
    wg.Wait()	
}

//------------------------------------------------------------------------
