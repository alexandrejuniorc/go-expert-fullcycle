package main

func soma(a, b *int) int {
	*a = 50
	*b = 50
	return *a + *b
}

func main() {
	minhaVar1 := 10
	minhaVar2 := 20

	// soma(minhaVar1, minhaVar2) // Passagem por valor (cópia dos valores)
	// println(minhaVar1)           // Imprime 10, pois minhaVar1 não foi alterada

	soma(&minhaVar1, &minhaVar2) // Passagem por referência (endereço dos valores)
	println(minhaVar1)           // Imprime 50, pois minhaVar1 foi alterada dentro da função
	println(minhaVar2)           // Imprime 50, pois minhaVar2 foi alterada dentro da função
}
