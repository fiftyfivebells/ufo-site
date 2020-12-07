package main

import (
	"flag"
	"log"
	"net/http"
)

func main() {
	addr := flag.String("addr", ":3000", "HTTP network address")
	flag.Parse()

	mux := http.NewServeMux()
	mux.HandleFunc("/", home)

	log.Printf("Starting server on %s", *addr)
	err := http.ListenAndServe(*addr, mux)
	log.Fatal(err)
}
