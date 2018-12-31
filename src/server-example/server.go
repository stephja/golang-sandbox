package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {

	// http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	// 	fmt.Fprintf(w, "index", html.EscapeString(r.URL.Path))
	// 	http.ServeFile(w, r, r.URL.Path[1:])
	// })

	http.Handle("/", http.FileServer(http.Dir("./static")))

	http.HandleFunc("/users", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "users")
	})

	log.Fatal(http.ListenAndServe(":8000", nil))

}
