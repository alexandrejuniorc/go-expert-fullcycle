package main

import "net/http"

func main() {

	// Define a rota "/" e associa à função BuscaCEP
	http.HandleFunc("/", BuscaCEP)

	// Inicia um servidor HTTP na porta 8080
	http.ListenAndServe(":8080", nil)
}

// BuscaCEP responde às requisições na rota "/"
// Funcionalidade separada conhecida também como controller ou handler
func BuscaCEP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello, World!"))
}
