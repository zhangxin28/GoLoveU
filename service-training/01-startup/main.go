package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	// Convert the Echo function to a type that implments http.Handler
	h := http.HandlerFunc(Echo)

	// Start a server listening on port 8000 and responding using Echo.
	if err := http.ListenAndServe("localhost:8000", h); err != nil {
		log.Fatalf("error: listening and serving: %s", err)
	}
}

// Echo is a basic Http Handler
func Echo(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "You asked to %s %s \n", r.Method, r.URL.Path)
}
