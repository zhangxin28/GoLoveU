package main

import (
	"log"
	"net/http"
)

/*
The Go language supports `first-class functions`; that is,
functions in Go are treated as "first-class citizens".
Functions can be assigned to variables, and then called from them.
We can pass functions to other functions.
We can use functions as types.
A function's parameters and return value are part of its type.
A variable that holds a function needs to specify what parameters and
return values that function should ahve.
The variable can only hold functions whose number and types of
parameters and return values match the specified type.
*/

func write(writer http.ResponseWriter, message string) {
	if _, err := writer.Write([]byte(message)); err != nil {
		log.Fatal(err)

	}
}

func viewHandler(writer http.ResponseWriter, request *http.Request) {
	write(writer, "hello, web")
}

// RunHello represents a simple web server
// when you type `http://localhost:8080/hello`
// you can get the response `Hello, web` on the browser
func RunHello() {
	http.HandleFunc("/hello", viewHandler)
	http.HandleFunc("/english", func(writer http.ResponseWriter, request *http.Request) {
		write(writer, "hello, english")
	})

	err := http.ListenAndServe("localhost:8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}
