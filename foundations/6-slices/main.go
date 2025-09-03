package main

import "fmt"

func main() {
	slice := []int{10, 20, 30, 40, 50, 60, 70, 80, 90, 100} // Declarando um slice com 10 elementos
	fmt.Printf("length=%d capacity=%d content=%v\n", len(slice), cap(slice), slice)

	// No final das contas, sempre que eu uso um [:indice], eu estou criando um novo slice que aponta para o mesmo array.

	// Quando eu utilizo o [:indice] eu estou pegando do início até o índice - 1
	// [:0] é um slice vazio
	fmt.Printf("length=%d capacity=%d content=%v\n", len(slice[:0]), cap(slice[:0]), slice[:0]) // Exibindo informações do slice vazio

	// [4:] é um slice com 4 elementos, do indice 0 até 3
	fmt.Printf("length=%d capacity=%d content=%v\n", len(slice[:4]), cap(slice[:4]), slice[:4]) // Exibindo informações do slice com 4 elementos

	// Quando eu utilizo o [indice:] eu estou pegando do índice até o final
	// [2:] é um slice com 8 elementos, ignorando os 2 primeiros indices
	fmt.Printf("length=%d capacity=%d content=%v\n", len(slice[2:]), cap(slice[2:]), slice[2:]) // Exibindo informações do slice com 8 elementos

	// Como aumentar a capacidade do slice?
	slice = append(slice, 110)                                                                  // Adicionando mais um elemento ao slice
	fmt.Printf("length=%d capacity=%d content=%v\n", len(slice[2:]), cap(slice[2:]), slice[2:]) // Exibindo informações do slice após o append
	fmt.Printf("length=%d capacity=%d content=%v\n", len(slice), cap(slice), slice)             // Exibindo informações do slice após o append
	// Porque a capacidade foi de 10 para 20?
	// Porque o slice original tinha capacidade para 10 elementos, mas ao adicionar um novo elemento,
	// o Go automaticamente alocou um novo array com o dobro da capacidade.
}
