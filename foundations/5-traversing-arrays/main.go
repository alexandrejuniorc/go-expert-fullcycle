package main

import "fmt"

type ID int

func main() {
	var meuArray [3]int // Declarando e inicializando um array com 3 posições
	meuArray[0] = 10    // Atribuindo o valor 10 à primeira posição
	meuArray[1] = 20    // Atribuindo o valor 20 à segunda posição
	meuArray[2] = 30    // Atribuindo o valor 30 à terceira posição

	fmt.Println(len(meuArray) - 1) // Retorna o índice da última posição
	fmt.Println(len(meuArray))     // Retorna o tamanho do array

	// Percorrendo o array com um loop for
	// range significa "intervalo" e é usado para iterar sobre elementos, ou percorrer uma coleção
	// (meuArray) é o array que estamos percorrendo
	for index, value := range meuArray {
		// quando for string é só colocar %v e se for digito é só usar %d
		fmt.Printf("O valor do índice %d é %d\n", index, value)
	}
}
