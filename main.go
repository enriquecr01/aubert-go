package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

func homeHandle(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello World!"))
}

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/", homeHandle)

	http.ListenAndServe(":3000", r)
}
