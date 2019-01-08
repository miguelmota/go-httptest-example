package main

import (
	"log"
	"net/http"

	"github.com/miguelmota/go-httptest-example/server"
)

func main() {
	mux := http.NewServeMux()
	handler := &server.Handler{}
	mux.Handle("/", handler)

	svr := http.Server{
		Handler: mux,
		Addr:    ":3333",
	}

	log.Printf("Listening on %s\n", svr.Addr)
	log.Fatal(svr.ListenAndServe())
}
