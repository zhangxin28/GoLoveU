package hexchan

import (
	"fmt"
	"starbucks-tools/utils"
)

var HexDataDispatch = make(chan utils.DispatchHandled)

type tests struct {
	Id   int
	Name string
}

var ch = make(chan tests)

func foo(id int) { //id: 这个routine的标号
	ch <- tests{id + 1, fmt.Sprintf("%di", id+1)}
}

func testChan() {
	// 开启5个routine
	for i := 0; i < 5; i++ {
		go foo(i)
	}

	// 取出信道中的数据
	for i := 0; i < 5; i++ {
		d := <-ch
		fmt.Printf("id=%d,name=%s\t", d.Id, d.Name)
	}

	sliceA := make([]int, 10, 10)
	fmt.Println(sliceA)
	fmt.Printf("%T\n", sliceA)
	sliceA = append(sliceA, 8)
	fmt.Println(sliceA)

	fmt.Println("Success")
}
