package main

import "fmt"

// Structs são tipos de dados compostos, que agrupam zero ou mais campos
// relacionados, formando um novo tipo de dado.
type Cliente struct {
	Nome  string
	Idade int
	Ativo bool
}

func main() {
	// Criando uma instância da struct Cliente
	alexandre := Cliente{
		Nome:  "Alexandre",
		Idade: 24,
		Ativo: true,
	}

	alexandre.Ativo = false // Alterando o valor do campo Ativo

	fmt.Println(alexandre.Nome)
	fmt.Println(alexandre.Idade)
	fmt.Println(alexandre.Ativo)
}
