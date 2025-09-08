package main

import "fmt"

type Endereco struct {
	Logradouro string
	Numero     int
	Cidade     string
	Estado     string
}

// Definindo uma interface
// A interface nos possibilita definir um contrato que tipos concretos
// devem seguir, garantindo que implementem certos métodos.
// Qualquer struct que implemente o método Desativar, está implicitamente
// satisfazendo a interface Pessoa.
type Pessoa interface {
	// A interface permite definir apenas métodos, não campos.
	Desativar() // Método que deve ser implementado por qualquer tipo que satisfaça a interface
}

type Empresa struct {
	Nome string
}

// Método Desativar associado ao tipo Empresa
// Note que o receptor do método é uma cópia do Empresa (valor).
// Portanto, alterações feitas em e dentro deste método não afetarão
// a instância original.
func (e Empresa) Desativar() {}

type Cliente struct {
	Nome     string
	Idade    int
	Ativo    bool
	Endereco Endereco
}

// Método Desativar associado ao tipo Cliente
// Note que o receptor do método é uma cópia do Cliente (valor).
// Portanto, alterações feitas em c dentro deste método não afetarão
// a instância original.
func (c Cliente) Desativar() {
	c.Ativo = false
	fmt.Printf("O cliente %s foi desativado\n", c.Nome)
}

// Função que aceita qualquer tipo que satisfaça a interface Pessoa
func Desativacao(pessoa Pessoa) {
	pessoa.Desativar()
}

func main() {

	alexandre := Cliente{
		Nome:  "Alexandre",
		Idade: 24,
		Ativo: true,
	}
	Desativacao(alexandre) // Passando o Cliente para a função que aceita a interface Pessoa

	minhaEmpresa := Empresa{
		Nome: "Minha Empresa",
	}
	Desativacao(minhaEmpresa) // Passando a Empresa para a função que aceita a interface Pessoa

}
