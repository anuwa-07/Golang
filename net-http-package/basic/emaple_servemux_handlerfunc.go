// ServMux have it own HandleFunc, which is do the both jobs,
// 			- Convert to Handler
//			- Assing to a specipic endpoint

package main

import (
	"fmt"
	"log"
	"net/http"
)

func simpleMessageHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World! Welcome to GO!")
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/home", simpleMessageHandler)

	log.Println("Listning . . . ")
	http.ListenAndServe(":8080", mux)
}