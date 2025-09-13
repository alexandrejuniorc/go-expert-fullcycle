package main

// Soma retorna a soma dos valores do map com tipo int
func SomaInteiro(m map[string]int) int {
	var soma int

	for _, v := range m {
		soma += v
	}

	return soma
}

// SomaFloat retorna a soma dos valores do map com tipo float64
func SomaFloat(m map[string]float64) float64 {
	var soma float64

	for _, v := range m {
		soma += v
	}

	return soma
}

type MyNumber int

// Constraint para tipos numéricos usando Generics
type Number interface {
	// ~ permite tipos que são baseados em int ou float64
	~int | float64
}

// SomaGenerics retorna a soma dos valores do map usando Generics
// T é um tipo genérico que pode ser int ou float64
// Number é uma interface que define os tipos permitidos
func SomaGenerics[T Number](m map[string]T) T {
	var soma T

	for _, v := range m {
		soma += v
	}

	return soma
}

// Compara compara dois valores de qualquer tipo que seja comparável
// T é um tipo genérico que deve ser comparável
// comparable é uma constraint embutida que permite tipos que suportam comparação
func Compara[T comparable](a, b T) bool {
	if a == b {
		return true
	}
	return false
}

func main() {
	m := map[string]int{"João": 10, "Maria": 20, "Pedro": 30}
	println(SomaInteiro(m))

	m2 := map[string]float64{"João": 10.5, "Maria": 20.3, "Pedro": 30.2}
	println(SomaFloat(m2))

	m3 := map[string]float64{"João": 10.5, "Maria": 20.3, "Pedro": 30.2}
	println(SomaGenerics(m3))

	m4 := map[string]MyNumber{"João": 10, "Maria": 20, "Pedro": 30}
	println(SomaGenerics(m4))

	println(Compara(10, 120))
}
