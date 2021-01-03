package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

/*
A different program might have to wait for user input.
And another might have to wait while data is read in from a file.
There are lots of situations where programs are just sitting around waiting.
`Concurrency` allows a program to pause one task and work on other tasks.
A program wating for user input might do other processing in the backgound.
A program might update a progress bar while reading from a file.
If a program is written to support concurrency,
then it may also support parallelism: running tasks simultaneously.
In Go, concurrent tasks are called `gorountines`.
Other programming languages have a similar concept called `threads`.
But `gorountines` require less memory, and less time to start up and stop.
`Goroutines` allow for concurrency: pausing one task to work on others.
And in some situations they allow parallelism:
working on multiple tasks simultaneously.
Use `go functioncall()` to start the goroutines.
Rememeber, the `go xxx` statement can't be used with return values.
Go won't let you use the return value from a function that called
with a go statement, because there is no guarantee the return value
will be ready before we attempt to use it.
But there is a way to communicate between goroutines: `channels`.
Not only do `channels` allow you to send values from one goroutine
to another, they ensure the sending goroutine has sent the value before
the receiving goroutine attemps to use it.
Sending and receiving values with `channels`.
create a channel: myChannal := make(chan float64)
send value: myChannel <- 3.14
receive value: a := <- myChannel
*/

func getPageSize(url string) (int, error) {
	fmt.Println("Getting", url)
	response, err := http.Get(url)
	if err != nil {
		return 0, err
	}

	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return 0, err
	}

	return len(body), nil
}

func responseSize(url string) {
	size, err := getPageSize(url)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(size)
}

type urlsize struct {
	url  string
	size int
}

func (u urlsize) String() string {
	return fmt.Sprintf("url \"%s\" has size of %d", u.url, u.size)
}

func responseSizeWithChannel(url string, channel chan urlsize) {
	size, err := getPageSize(url)
	if err != nil {
		panic(err)
	}

	channel <- urlsize{url, size}
}

// RunRetrievingPagesSync get pages content length one by one
func RunRetrievingPagesSync() {
	responseSize("https://www.qq.com")
	responseSize("https://www.163.com")
	responseSize("https://www.baidu.com")
	fmt.Println("Getting urls done")
}

// RunRetrievingPaegsUsingGoroutines gets pages content length via goroutines
// When call the function, we see nothing on the console.
// Because the main goroutine ends, no more enouth time for the other goroutines to run
func RunRetrievingPaegsUsingGoroutines() {
	go responseSize("https://www.qq.com")
	go responseSize("https://www.163.com")
	go responseSize("https://www.baidu.com")
	fmt.Println("Getting urls done")
}

// Below samples are using goroutines.
// And the order that calles to responseSize is not under our control.
// So if you run the program more times,
// you may see the requests happen in a different order.
// We don't directly control when gorountines run.

// RunRetrievingPagesUsingGoroutinesAndTimeSleep gets pages content
// length via goroutine, and use time.Sleep to delay the main goroutine
// to exit, and gives other goroutines more time to finish running.
// The function takes 5 seconds to complete even if all the sites
// respond faster than that, so we're still not getting that great
// a speed gain from the switch to goroutines.
// Even worse, 5 seconds may not be enough time if the sites takes
// a long time to respond.
// Sometimes, you may see the program end before all the responses have arrived.
func RunRetrievingPagesUsingGoroutinesAndTimeSleep() {
	go responseSize("https://www.qq.com")
	go responseSize("https://www.163.com")
	go responseSize("https://www.baidu.com")
	time.Sleep(5 * time.Second)
	fmt.Println("Getting urls done")
	runSimpleChannel()
}

// RunRetrievingPagesUsingChannel gets page content size using channel
func RunRetrievingPagesUsingChannel() {
	defer func() {
		if r := recover(); r != nil {
			log.Fatal(r)
		}
	}()

	sizes := make(chan urlsize)
	urls := []string{
		"https://www.qq.com",
		"https://www.163.com",
		"https://www.baidu.com",
	}
	for _, url := range urls {
		go responseSizeWithChannel(url, sizes)
	}

	for i := 0; i < len(urls); i++ {
		fmt.Println(<-sizes)
	}
}

func runSimpleChannel() {
	myChannel := make(chan string)

	go func(channel chan string) {
		channel <- "hi"
	}(myChannel)

	fmt.Println(<-myChannel)

	channel1 := make(chan string)
	go func(channel chan string) {
		channel <- "a"
		channel <- "b"
		channel <- "c"
	}(channel1)

	channel2 := make(chan string)
	go func(channel chan string) {
		channel <- "d"
		channel <- "e"
		channel <- "f"
	}(channel2)

	// output is : adbecf
	fmt.Print(<-channel1)
	fmt.Print(<-channel2)
	fmt.Print(<-channel1)
	fmt.Print(<-channel2)
	fmt.Print(<-channel1)
	fmt.Print(<-channel2)
	fmt.Println()

	reportNap := func(name string, delay int) {
		for i := 0; i < delay; i++ {
			fmt.Println(name, "sleeping", i)
			time.Sleep(1 * time.Second)
		}
		fmt.Println(name, "wakes up!")
	}
	channel3 := make(chan string)
	go func(channel chan string) {
		reportNap("sending goroutine", 2)
		fmt.Println("***sending value***")
		channel <- "a"
		fmt.Println("***sending value***")
		channel <- "b"
	}(channel3)
	reportNap("receiving goroutine", 5)
	fmt.Println(<-channel3)
	fmt.Println(<-channel3)
}
