package main

import (
	"fmt"
	"log"
	"net/http"
)

// Add a Norml func as http.Handler func
func customeReqHandler01(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello! Welcome to Golang!")
}


// using Closure
func messageHandler(message string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, message)
	})
}



func main() {
	// Define a Serve to handle the response
	mux := http.NewServeMux()

	// convert customeReqHandler01 func to a HandlerFunc
	cutmReqHdler := http.HandlerFunc(customeReqHandler01)
	mux.Handle("/welcome", cutmReqHdler)
	mux.Handle("/message", messageHandler("net/http is awesome"))

	log.Println("Listening . . . ")
	http.ListenAndServe(":8080", mux)

}










