
# NET/HTTP package
---

> Have Two Majro components
    - ServeMux
    - Handler


> ServerMux
    - The ServeMux is a multiplexor (or simply an HTTP request router) that compares incoming HTTP requests against a list of predefined URI resources and then calls the associated handler for the resource requested by the HTTP client.

> Handler
    - Handlers are responsible for writing response headers and bodies.
    -  If any object satisfies the implementation of the http.Handler interface, it can be a handler for serving HTTP requests.

    ```
    type Handler interface {
        ServeHTTP(ResponseWriter, *Request)
    }
    ```

Example Code:

```
package main

import (
    "fmt"
    "log"
    "net/http"
)

func messageHandler(w http.ResopnseWritter, r *http.Request){
    fmt.Fprintf(w, "Welcome to the GO Web Development")
}

func main() {
    mux = http.NewServeMux()

    // Convert messageHandler func to a HandlerFunc
    mh := http.HandlerFunc(messaheHandler)
    mux.Handle("/welcome", mh)
    
    log.Println("Listening. . . ")
    http.ListenAndServe(":8000", mux)
}

```
---
