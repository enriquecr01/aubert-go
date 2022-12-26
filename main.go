package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func homeHandle(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello me llamo Aubert!"))
}

func exampleHandle(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Salut ç'est un exemple"))
}

func main() {
	fmt.Println("Server initialized")

	r := mux.NewRouter()

	r.HandleFunc("/", homeHandle)
	r.HandleFunc("/example", exampleHandle)

	http.ListenAndServe(":80", r)
}
