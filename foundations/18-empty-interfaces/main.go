package main

import "fmt"

func main() {

	// Interface vazia implementa qualquer tipo, ou seja, qualquer tipo é um subconjunto do tipo interface{}.
	// Ela é usada quando queremos trabalhar com valores de tipos desconhecidos.
	var x interface{} = 10
	var y interface{} = "Hello, World!"

	showType(x)
	showType(y)
}

func showType(t interface{}) {
	fmt.Printf("O tipo da variável é %T e o valor é %v\n", t, t)
}
