package main

import (
	"flag"
	"log"
	"net/http"
	"strconv"
)

type server struct {
	router http.Handler
}

func newServer() *server {
	s := &server{}
	return s
}

func main() {
	// Flag to read port
	port := flag.Int("port", 8080, "Server listens on this port")
	flag.Parse()

	s := newServer()

	s.router = initialiseRoutes()

	log.Printf("Server Listening on %v", *port)
	log.Fatal(http.ListenAndServe(":"+strconv.Itoa(*port), s.router))
}
