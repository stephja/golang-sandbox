package main

import (
    "fmt"
    "html"
    "log"
    "net/http"
)

func main() {

    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
    })
    
    http.HandleFunc("/users", func(w http.ResponseWriter, r *http.Request){
        fmt.Fprintf(w, "users")
    })

    log.Fatal(http.ListenAndServe(":8000", nil))

}