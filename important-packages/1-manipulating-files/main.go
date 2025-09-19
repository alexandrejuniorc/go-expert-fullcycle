package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// Cria um arquivo chamado "test.txt"
	file, err := os.Create("test.txt")
	if err != nil {
		panic(err)
	}

	// Escreve um slice de bytes no arquivo
	tamanho, err := file.Write([]byte("Escrevendo dados no arquivo"))
	// Escreve uma string no arquivo
	// tamanho, err := file.WriteString("Hello, Go!")
	if err != nil {
		panic(err)
	}
	fmt.Printf("Arquivo criado com sucesso! Tamanho: %d bytes\n", tamanho)
	file.Close()

	// Lê o conteúdo do arquivo "test.txt"
	arquivo, err := os.ReadFile("test.txt")
	if err != nil {
		panic(err)
	}
	// Imprime o conteúdo do arquivo
	// Está em formato de slice de bytes, então convertemos para string
	fmt.Println(string(arquivo))

	// Leitura de pouco em pouco abrindo o arquivo
	arquivo2, err := os.Open("test.txt")
	if err != nil {
		panic(err)
	}
	// Cria um buffer para ler o arquivo
	reader := bufio.NewReader(arquivo2) // Leitor com buffer
	buffer := make([]byte, 3)           // Buffer de 10 bytes

	// Lê o arquivo pouco a pouco, de 10 em 10 bytes
	for {
		// Lê até preencher o buffer ou encontrar um erro
		n, err := reader.Read(buffer)
		if err != nil {
			break
		}
		// Imprime o conteúdo lido
		fmt.Println(string(buffer[:n]))
	}

	// Remove o arquivo "test.txt"
	err = os.Remove("test.txt")
	if err != nil {
		panic(err)
	}
	fmt.Println("Arquivo removido com sucesso!")
}
