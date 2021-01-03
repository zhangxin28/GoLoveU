package main

import (
	"fmt"
	"io/ioutil"
	"path/filepath"
)

/*
`recursion` is a function which can call itself.
If you write a `recursion` function carelessly, you'll just
wind up with an infinite loop where the function calls itself
over and over again, and never stops.
eg:

func recurses() {
	fmt.Println("Oh, no, I'm stuck!")
	recurses()
}

The `defer` key word can be added before any function or method
call to postpone that call until the current function exists.
Deferred function calls are often used for cleanup code that needs
to be run even in the event of an error.

If a deferred function calls the build-in `recover` function,
the program will recover from a panic state.
The `recover` function returns whatever value was originally
passed to the `panic` function.
*/

func count(start int, end int) {
	fmt.Printf("count(%d,%d) called\n", start, end)
	fmt.Println(start)
	if start < end {
		count(start+1, end)
	}
	fmt.Printf("returning from count(%d,%d) call\n", start, end)
}

func reportPanic() {
	fmt.Println("Now we are doing panic/recover")
	p := recover()
	if p == nil {
		return
	}

	err, ok := p.(error)
	if ok {
		fmt.Println(err)
	} else {
		panic(p)
	}
}

func scanDirectory(path string) error {
	fmt.Println(path)
	files, err := ioutil.ReadDir(path)
	if err != nil {
		fmt.Printf("returning error from scanDirectory(\"%s\") called\n",
			path)
		return err
	}

	for _, file := range files {
		filePath := filepath.Join(path, file.Name())
		if file.IsDir() {
			err := scanDirectory(filePath)
			if err != nil {
				fmt.Printf("returning error from scanDirectory(\"%s\") call\n",
					path)
				return err
			}
		} else {
			fmt.Println(filePath)
		}
	}

	return nil
}

func scanDirectoryWithPanic(path string) {
	fmt.Println(path)
	files, err := ioutil.ReadDir(path)
	if err != nil {
		fmt.Printf("returning error from scanDirectory(\"%s\") called\n",
			path)
		panic(err)
	}

	for _, file := range files {
		filePath := filepath.Join(path, file.Name())
		if file.IsDir() {
			scanDirectoryWithPanic(filePath)
			fmt.Printf("returning error from scanDirectory(\"%s\") call\n",
				path)
		} else {
			fmt.Println(filePath)
		}
	}
}

// RunFiles tests to read files and directories in a folder
func RunFiles() {
	defer reportPanic()
	// for testing  error in `recursion` we set the inner
	// folder `locked` unaccessable
	err := scanDirectory("my_directory")
	if err != nil {
		//log.Fatal(err)
		panic(err)
	}
	count(1, 5)
}

// RunFilesWithPanic tests the panic/recover/recursion
func RunFilesWithPanic() {
	defer reportPanic()

	a := func() (b int) {
		defer func() {
			if err := recover(); err != nil {
				fmt.Println(err)
				b = 1
			}
		}()
		_, err := ioutil.ReadDir("my_directory/locked")
		if err != nil {
			panic(err)
		}
		return 0
	}()
	fmt.Println(a)

	// panic("other panic not error")
	scanDirectoryWithPanic("my_directory")
	count(1, 5)
}

// RunDeferWithoutReturnAndPanic tests the defer function run without return and panic
func RunDeferWithoutReturnAndPanic() {
	defer fmt.Println("run after the last")
	defer reportPanic()

	fmt.Println("Read a folder")
	_, err := ioutil.ReadDir("my_directory/locked")
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Now I am the last code")
}
