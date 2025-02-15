package main

import (
	"io"
	"log"
	"net/http"
	"time"
)

func main() {
	// Handles all traffic '/' to `servePage` function
	http.HandleFunc("/hello", servePage)

	// To serve static files from server
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, r.URL.Path[1:])
	})

	// To start the server. ListenAndServe(port as string, HTTP request multiplexer aka ServeMux)
	// If it is nil, the default ServeMux is used: DefaultServMux)
	// http.ListenAndServe(":8080", nil)
	// Can also do the following
	log.Fatal(http.ListenAndServe(":8080", nil))

	// In order to serve content over HTTPS we can use the following
	// First param is the HTTPS port, next is the server cert, server key and multiplexer
	log.Fatal(http.ListenAndServeTLS(":443", "server.crt", "server.key", nil))
}

// w allows to write into the response, r is a pointer to access the Request
func servePage(w http.ResponseWriter, r *http.Request) {
	// To write into the response
	io.WriteString(w, "Hello \nThe current time is:"+time.Now().String())
}
