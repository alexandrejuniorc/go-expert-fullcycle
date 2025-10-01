package main

import (
	"bytes"
	"io"
	"net/http"
	"os"
)

func main() {
	c := http.Client{}
	// Exemplo de POST com JSON
	jsonVar := bytes.NewBuffer([]byte(`{"name": "Alexandre"}`))
	// Fazendo a requisição POST passando parametros de URL, Content-Type e Body
	resp, err := c.Post("https://google.com", "application/json", jsonVar)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	// Imprimindo o Body da resposta no console
	io.CopyBuffer(os.Stdout, resp.Body, nil)
}
