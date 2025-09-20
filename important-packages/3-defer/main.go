package main

import (
	"fmt"
	"net/http"
)

func main() {
	request, err := http.Get("https://www.google.com")
	if err != nil {
		panic(err)
	}

	// O defer é usado para garantir que o recurso seja fechado
	// mesmo que ocorra um panic ou erro na função
	defer request.Body.Close() // Fecha o corpo da resposta ao final da função

	defer fmt.Println("Primeira Linha") // Será executado no final da função main
	fmt.Println("Segunda Linha")        // Será executado antes do fechamento do corpo da resposta
	fmt.Println("Terceira Linha")       // Será executado antes do fechamento do corpo da resposta

}
