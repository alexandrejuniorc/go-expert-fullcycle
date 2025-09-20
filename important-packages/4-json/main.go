package main

import (
	"encoding/json"
	"os"
)

type Conta struct {
	// Para mapear os nomes dos campos no JSON, usamos tags
	// Caso eu não queira disponibilizar um campo, basta começar com letra minúscula, basta coloca - na tag
	// Ex: Saldo int `json:"-"`
	Numero int `json:"numero"` // O nome entre aspas é o nome que vai aparecer no JSON
	Saldo  int `json:"saldo"`  // Se não colocar a tag, o nome do campo no JSON será o mesmo do struct
}

func main() {
	conta := Conta{Numero: 1, Saldo: 1000}
	res, err := json.Marshal(conta)
	if err != nil {
		panic(err)
	}
	println(res)         // Retorna em bytes
	println(string(res)) // Retorna em json

	err = json.NewEncoder(os.Stdout).Encode(conta) // Escreve diretamente no output
	if err != nil {
		panic(err)
	}

	jsonPuro := []byte(`{"numero":2,"saldo":2000}`) // JSON "puro", que pode vir de qualquer lugar
	var contaX Conta
	err = json.Unmarshal(jsonPuro, &contaX) // Faz o inverso do Marshal, apontando para a variável onde quer guardar usando ponteiro "&"
	if err != nil {
		panic(err)
	}
	println(contaX.Saldo) // 2000
}
