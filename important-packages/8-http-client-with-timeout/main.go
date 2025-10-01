package main

import (
	"io"
	"net/http"
	"time"
)

func main() {
	// Timeout de 1 microssegundo para forçar o timeout
	// Timeout de 1 segundo para permitir a resposta
	// Timeout de 10 segundos para permitir a resposta
	c := http.Client{Timeout: time.Microsecond}
	resp, err := c.Get("https://google.com")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	println(string(body))
}
