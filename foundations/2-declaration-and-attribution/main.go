package main

// Variáveis de escopo global
const a = "Hello, World!"

var (
	b bool = true // Declarando e atribuindo valor
	c int
	d string
	e float64
) // Declaração de mais de uma variável ao mesmo tempo

func main() {
	// := é uma forma curta de declarar e inicializar variáveis
	// e é utilizada pra declarar a variável pela primeira vez
	shortVariable := "X" // Declaração e atribuição de variável de uma forma mais curta
	println(shortVariable)

	var a string // Declaração de variável local
	println(a)   // Imprime o valor padrão (string vazia)

	b = true // Atribuição de valor à variável booleana
	// b = 10 -> Erro de compilação: cannot use 10 (type int) as type bool in assignment
	println(b)

	c = 42 // Atribuição de valor à variável inteira
	println(c)

	d = "Hello, Go!" // Atribuição de valor à variável string
	println(d)

	e = 3.14 // Atribuição de valor à variável float64
	println(e)
}
