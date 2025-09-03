package main

import (
	"errors"
	"fmt"
)

// Função principal
func main() {
	// fmt.Println(sum(49, 2)) // Chama a função sum e imprime o resultado

	valor, err := sum(1, 2) // Chama a função sum e trata o erro

	// Verifica se houve erro
	if err != nil {
		fmt.Println(err)
	}

	// Imprime o valor retornado se não houver erro, mas se houver irá retornar 0
	fmt.Println(valor)

}

// Função para somar dois números que retorna a soma
// func sum(a int, b int) int {
// 	return a + b
// }

// Quando tenho dois valores que são do mesmo tipo, posso omitir o tipo na primeira variável
// Posso retornar mais de um valor, por exemplo, um resultado e um status
func sum(a, b int) (int, error) {
	if a+b >= 50 {
		return 0, errors.New("A soma é maior ou igual que 50")
	}

	return a + b, nil
}
