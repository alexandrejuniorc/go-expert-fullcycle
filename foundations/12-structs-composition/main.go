package main

import "fmt"

type Endereco struct {
	Logradouro string
	Numero     int
	Cidade     string
	Estado     string
}

type Cliente struct {
	Nome  string
	Idade int
	Ativo bool
	// Endereco // Composição de structs, onde Cliente "tem um" Endereco
	Endereco Endereco // Outra forma de declarar a composição de structs, criando um campo chamado Endereco do tipo Endereco
}

func main() {

	alexandre := Cliente{
		Nome:  "Alexandre",
		Idade: 24,
		Ativo: true,
	}

	alexandre.Ativo = false
	alexandre.Endereco.Cidade = "São Paulo"   // Acessando o campo da struct Endereco diretamente
	alexandre.Endereco.Logradouro = "Rua ABC" // Outra forma de acessar o campo da struct Endereco

	fmt.Println(alexandre)

}
