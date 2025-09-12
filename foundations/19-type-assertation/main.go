package main

import "fmt"

func main() {
	var minhaVar interface{} = "Alexandre Junior"

	println(minhaVar)          // Não tem tipo definido
	println(minhaVar.(string)) // Tem tipo definido

	res, ok := minhaVar.(int)                                             // Verifica se é do tipo int
	fmt.Printf("O valor de res é %v e o resultado de ok é %v\n", res, ok) // O valor de res é 0 e o resultado de ok é false

	// Irá dar panic, pois, não teve validação do ok antes da asserção
	res2 := minhaVar.(int)
	fmt.Printf("O valor de res2 é %v\n", res2) // Panic: interface conversion: interface {} is string, not int
}
