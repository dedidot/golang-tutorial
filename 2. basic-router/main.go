package main

import (
	"fmt"
	"net/http"
	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", handler).Methods("GET")
	r.HandleFunc("/hello", helloHandle).Methods("GET")
	http.Handle("/", r)
	http.ListenAndServe(":8000", nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello world")
}

func helloHandle(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Oke ini hello")
}