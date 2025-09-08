package main

import "fmt"

type Endereco struct {
	Logradouro string
	Numero     int
	Cidade     string
	Estado     string
}

type Cliente struct {
	Nome     string
	Idade    int
	Ativo    bool
	Endereco Endereco
}

func (c Cliente) Desativar() {
	c.Ativo = false
	fmt.Printf("O cliente %s foi desativado\n", c.Nome)
}

func main() {

	alexandre := Cliente{
		Nome:  "Alexandre",
		Idade: 24,
		Ativo: true,
	}

	alexandre.Desativar()

}
