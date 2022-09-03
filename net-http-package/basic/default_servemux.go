
// Here we are using the default serveMux, So we do not need to make a new One!

package main

import (
	"fmt"
	"log"
	"net/http"
)

func myMessageHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the Go web development!")
}

func main() {
	http.HandleFunc("/welcome", myMessageHandler)
	log.Println("Listen on Port 5000")
	http.ListenAndServe(":5000", nil)
}


