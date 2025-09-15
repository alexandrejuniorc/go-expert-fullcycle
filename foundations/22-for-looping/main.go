package main

func main() {
	// Estrutura de repetição mais utilizada em Go
	// for init; condition; post { }
	// init -> inicialização da variável de controle
	// condition -> condição de continuação do loop
	// post -> incremento ou decremento da variável de controle

	// Exemplo:
	for i := 0; i < 10; i++ {
		println(i)
	}

	numeros := []string{"um", "dois", "três", "quatro", "cinco"}
	// for range -> estrutura de repetição para iterar sobre arrays, slices, maps, strings e canais

	// Exemplo iterando sobre um slice
	// Posso também ignorar o índice ou o valor utilizando o underline _
	for key, value := range numeros {
		println(key, value)
	}

	// Outra forma de utilizar o for é como um while
	// for condition { }
	// Exemplo:
	i := 0
	for i < 10 {
		println(i)
		i++
	}

	// Loop infinito
	for {
		println("loop infinito")
		break // para evitar loop infinito, utilize break para sair do loop
	}
}
