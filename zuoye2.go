package main

import (
	"log"
	"net/http"
	"strings"
)

func main() {
	http.HandleFunc("/", rootHandler)
	http.HandleFunc("/healthz", healthzHandler)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}

func healthzHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("/healthz ClientAddress is " + r.RemoteAddr + " Status Code is " + http.StatusText(200))
}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	for k, v := range r.Header {
		w.Header().Add(k, strings.Join(v, " "))
	}
	log.Println("/ ClientAddress is " + r.RemoteAddr + " Status Code is " + http.StatusText(200))
}
