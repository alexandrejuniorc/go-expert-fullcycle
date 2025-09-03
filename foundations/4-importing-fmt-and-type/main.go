package main

import "fmt"

const a = "Hello, World!"

type ID int

var (
	b bool    = true
	c int     = 42
	d string  = "Hello, Go!"
	e float64 = 3.14
	f ID      = 1
)

func main() {

	// O Printf é uma função da biblioteca fmt que permite formatar e imprimir strings
	// Ela utiliza verbos de formato para especificar como os valores devem ser apresentados.
	fmt.Printf("O tipo de E é %T\n", e)  // %T traz o tipo da variável
	fmt.Printf("O valor de E é %v\n", e) // %v traz o valor da variável

	fmt.Printf("O tipo de F é %T\n", f)  // %T traz o tipo da variável e irá retornar o tipo personalizado
	fmt.Printf("O valor de F é %v\n", f) // %v traz o valor da variável
}
