package main

import (
	"fmt"

	"github.com/alexandrejuniorc/go-expert-fullcycle/matematica"
	"github.com/google/uuid"
)

func main() {
	// Usando a função Soma do pacote matematica
	soma := matematica.Soma(10, 20)

	fmt.Println("Resultado: ", soma)

	carro := matematica.Carro{Marca: "Toyota"}
	fmt.Println("Carro: ", carro.Marca)

	carro.Andar()

	fmt.Println(uuid.New())
}
