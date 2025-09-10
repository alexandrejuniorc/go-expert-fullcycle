package main

func main() {
	// Memória -> Endereço -> Valor
	// Variável -> ponteiro que tem um endereço na memória -> valor armazenado nesse endereço
	a := 10
	println(&a) // Endereço de memória da variável 'a'

	var ponteiro *int = &a // Ponteiro que armazena o endereço de 'a'
	println(ponteiro)      // Imprime o endereço armazenado no ponteiro
	println(*ponteiro)     // Desreferencia o ponteiro para obter o valor de 'a'
	*ponteiro = 20         // Altera o valor de 'a' através do ponteiro
	println(a)             // Imprime o novo valor de 'a'

	b := &a // Sintaxe curta para criar um ponteiro
	println(b)
	println(*b)
	*b = 30
	println(a)
}
