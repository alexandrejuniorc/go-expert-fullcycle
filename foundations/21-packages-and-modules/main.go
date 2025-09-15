package main

import (
	"fmt"

	"github.com/alexandrejuniorc/go-expert-fullcycle/matematica"
)

func main() {
	// Usando a função Soma do pacote matematica
	soma := matematica.Soma(10, 20)

	fmt.Println("Resultado: ", soma)

	carro := matematica.Carro{Marca: "Toyota"}
	fmt.Println("Carro: ", carro.Marca)

	carro.Andar()
}
