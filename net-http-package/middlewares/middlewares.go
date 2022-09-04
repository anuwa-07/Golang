/*

	HTTP Middlewares
		- Just a simple function act as a decorator.
		- Middleware func accepet parameter which type it http.Handler and also it return a http.Handler type func.
		- So that servMux will able to register that func with a perticular endpoint.
		- We can use "ServeHTTP(res, req)" to call handler func inside of Middleware

		Ex:
		func middlewareHandler(next http.Handler) http.Handler {
			return http.HandlerFunc( func(w http.ResponseWriter, r *http.Request) {
				
				// Middleware logic which need to apply before req pass to handler func will add here

				next.ServeHTTP(w, r)

				// Middleware logic goes here after executing application handler
			}
		}
*/

// Exaple of Simple Logging Middleware
package main
import (
	"fmt"
	"net/http"
	"time"
	"log"
)

// Add the middleware
func loggingHandler(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// define the middleware logic before call on http handler func
		start := time.Now()
		log.Printf("Started %s %s", r.Method, r.URL.Path)
		
		next.ServeHTTP(w, r)

		// define the logic aftert the http handler func call
		log.Printf("Completed %s in %v", r.URL.Path, time.Since(start))
	})
}

// Http Middlerware accept and return a http.Handler type func. So due to this, we can call combine multiple middlewares to some job done.
func middlewareSecond(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println("MiddlewareSecond - Before Handler")

		if r.URL.Path == "/message" {
			if r.URL.Query().Get("password") == "password@123" {
                log.Println("Authorized to the system")
                next.ServeHTTP(w, r)				
			} else {
                log.Println("Failed to authorize to the system")
                return				
			}
		}else {
			next.ServeHTTP(w, r)
		}

		log.Println("MiddlewareSecond - After Handler")
	})
}



func index(w http.ResponseWriter, r *http.Request) {
	log.Println("Executing index handler")
	fmt.Fprintf(w, "Welcome!")
}

func about(w http.ResponseWriter, r *http.Request) {
	log.Println("Executing about handler")
	fmt.Fprintf(w, "Go Middleware")
}

func message(w http.ResponseWriter, r *http.Request) {
    log.Println("Executing message Handler")
    fmt.Fprintf(w, "HTTP Middleware is awesome")
}


func iconHandler(w http.ResponseWriter, r *http.Request) {
}


func main() {
	http.HandleFunc("/favicon.ico", iconHandler)
	indexHandler := http.HandlerFunc(index)
	aboutHandler := http.HandlerFunc(about)
	messageHandler := http.HandlerFunc(message)

	http.Handle("/", loggingHandler(indexHandler))
	http.Handle("/about", loggingHandler(aboutHandler))
	http.Handle("/message", middlewareSecond(loggingHandler(messageHandler))) // multiple middleware usecases


	server := &http.Server{
		Addr: ":8080",
	}

	log.Println("Listening . . .")
	server.ListenAndServe()
}