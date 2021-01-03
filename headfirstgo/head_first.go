package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

type mystruct struct {
	name string
	age  int
}

func init() {
	fmt.Println("Current Go Version:", runtime.Version())
}

func main() {
	fmt.Println("Head First Go!")
	// RunPassFail()
	// RunGuessGame()
	// RunTocelsius()
	// RunReadFile()
	// RunCount()
	// RunCountWithMap()
	// RunMagazine()
	// RunReceiver()
	// RunCalendar()
	// RunMyInterface()
	// RunFiles()
	// RunFilesWithPanic()
	// RunDeferWithoutReturnAndPanic()
	// RunRetrievingPagesSync()
	// RunRetrievingPaegsUsingGoroutines()
	// RunRetrievingPagesUsingGoroutinesAndTimeSleep()
	RunRetrievingPagesUsingChannel()
	// RunHello()
	// runPool()
}

type pool struct {
	queue chan int
	wg    *sync.WaitGroup
}

func newPool(size int) *pool {
	if size < 0 {
		size = 1
	}
	return &pool{
		queue: make(chan int, size),
		wg:    &sync.WaitGroup{},
	}
}

func (p *pool) add(delta int) {
	for i := 0; i < delta; i++ {
		p.queue <- 1
	}

	for i := 0; i > delta; i-- {
		<-p.queue
	}
	p.wg.Add(delta)
}

func (p *pool) done() {
	<-p.queue
	p.wg.Done()
}

func (p *pool) wait() {
	p.wg.Wait()
}

func runPool() {
	runtime.GOMAXPROCS(8)
	pool := newPool(5)
	fmt.Println("the numgoroutine begin is:", runtime.NumGoroutine())
	for i := 0; i < 20; i++ {
		pool.add(1)
		go func(i int) {
			time.Sleep(time.Second)
			fmt.Println("the numgoroutine continue is:", runtime.NumGoroutine())
			pool.done()
		}(i)
	}
	pool.wait()
	fmt.Println("the numgorountine done is:", runtime.NumGoroutine())
}

func interestFunc1() {
	var x int
	inc := func() int {
		x++
		return x
	}
	fmt.Println(func() (a, b int) {
		return inc(), inc()
	}())
}

func interestFunc2() {

	c1 := make(chan int)
	c2 := make(chan int)

	go func() {
		c2 <- 1
	}()

	go func() {
		fmt.Println("11111111111")
		c1 <- <-c2
	}()

	go func() {
		c1 <- <-c2
		fmt.Println("22222222222")

	}()

	fmt.Println(len("“"))

	<-c1
}
