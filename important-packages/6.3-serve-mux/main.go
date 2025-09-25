package main

import (
	"net/http"
)

func main() {
	mux := http.NewServeMux()

	// Função anônima
	// mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	// 	w.Write([]byte("Hello, World!"))
	// })

	// Função anônima com handler dedicado
	mux.HandleFunc("/", HomeHandler)

	mux.Handle("/blog", Blog{title: "Blog do Alexandre"})

	// Mux com handler dedicado
	http.ListenAndServe(":8080", mux)
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello, World!"))
}

type Blog struct {
	title string
}

func (b Blog) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(b.title))
}
