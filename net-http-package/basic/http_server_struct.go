package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

/*
type Server struct {
		Addr
		string
		Handler
		Handler
		ReadTimeout
		time.Duration
		WriteTimeout
		time.Duration
		MaxHeaderBytes int
		TLSConfig
		*tls.Config
		TLSNextProto
		map[string]func(*Server, *tls.Conn, Handler)
		ConnState
		func(net.Conn, ConnState)
		ErrorLog
		*log.Logger
}
*/



func messageHandler02(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Welcome to GO Web Development!")
}

func main() {
	http.HandleFunc("/welcome", messageHandler02)

	// Here we Define the all configurations we need!
	server := &http.Server{
		Addr: ":8080",
		ReadTimeout: 10 * time.Second,
		WriteTimeout: 10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	log.Println("Listening . . . ")
	server.ListenAndServe()
}



