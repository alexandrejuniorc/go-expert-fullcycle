package matematica

import "fmt"

// Funções que começam com letra maiúscula são exportadas (públicas)
// Funções que começam com letra minúscula são não exportadas (privadas)

// função exportada
func Soma[T int | float64](a, b T) T {
	return a + b
}

// função não exportada
func soma[T int | float64](a, b T) T {
	return a + b
}

type Carro struct {
	Marca string
}

func (c Carro) Andar() {
	fmt.Println("O carro", c.Marca, "está andando.")
}
