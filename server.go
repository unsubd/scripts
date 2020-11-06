/*
A Simple file server written in Go
*/
package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
)

func main() {
	var port string
	flag.StringVar(&port, "p", "8100", "port to serve on")
	var directory string
	flag.StringVar(&directory, "d", ".", "the directory of static file to host")
	flag.Parse()

	http.Handle("/", http.FileServer(http.Dir(directory)))
	fmt.Printf("Serving %s on HTTP port: %s\n", directory, port)

	log.Fatal(http.ListenAndServe(":"+port, nil))
}
