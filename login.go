package main

import (
	"flag"
	"log"
	"net/http"
)

var addr, port *string

func init() {
	addr = flag.String("addr", "", "Server address")
	port = flag.String("port", "8080", "Server port")
}

type handler func(http.ResponseWriter, *http.Request)

func logger(h handler) handler {
	return func(w http.ResponseWriter, req *http.Request) {
		log.Println(req.Method, req.URL)
		h(w, req)
	}
}

func serveHttp(w http.ResponseWriter, req *http.Request) {}

func main() {
	flag.Parse()
	http.HandleFunc("/", logger(serveHttp))
	err := http.ListenAndServe(*addr+":"+*port, nil)
	if err != nil {
		log.Fatal(err)
	}
}
