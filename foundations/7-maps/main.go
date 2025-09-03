package main

import "fmt"

func main() {
	// Estrutura para o map é map[chave]valor
	salarios := map[string]int{"Alexandre": 1000, "Maria": 1500, "João": 2000} // Map de salários

	fmt.Println(salarios["Alexandre"]) // Exibindo o salário de Alexandre

	delete(salarios, "Alexandre") // Removendo o salário de Alexandre

	salarios["Ale"] = 5000 // Adicionando um novo salário

	fmt.Println(salarios["Ale"]) // Exibindo o salário de Ale

	// O make é utilizado para inicializar um map
	// salario := make(map[string]int) // Criando um map vazio
	// salario1 := map[string]int{}    // Outra forma de criar um map vazio
	// salario1["Alexandre"] = 1000    // Adicionando um novo salário

	for nome, salario := range salarios {
		fmt.Printf("O salário de %s é R$%d\n", nome, salario)
	}

	// Caso eu não queira usar uma variável de salários, ele implementa um blank identifier
	for _, salario := range salarios {
		fmt.Printf("O salário é R$%d\n", salario)
	}
}
