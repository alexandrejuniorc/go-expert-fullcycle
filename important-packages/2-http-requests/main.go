package main

import (
	"io"
	"net/http"
)

func main() {
	request, err := http.Get("https://www.google.com")
	if err != nil {
		panic(err)
	}

	response, err := io.ReadAll(request.Body) // Lê todo o conteúdo do corpo da resposta
	if err != nil {
		panic(err)
	}
	println(string(response))

	request.Body.Close() // Fecha stream de dados, pois, se não fechar, pode causar memory leak
	// Ou pode usar o defer para fechar automaticamente ao final da função main

}
