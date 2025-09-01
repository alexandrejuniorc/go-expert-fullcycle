package main

const a = "Hello, World!"

// No Go, podemos criar tipos personalizados
type ID int // Tipo personalizado para IDs

var (
	b bool    = true
	c int     = 42
	d string  = "Hello, Go!"
	e float64 = 3.14
	f ID      = 1 // Atribuindo valor com tipo personalizado
)

func main() {
	println(f)
}
