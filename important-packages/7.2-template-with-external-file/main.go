package main

import (
	"html/template"
	"os"
)

type Curso struct {
	Nome         string
	CargaHoraria int
}

type Cursos []Curso

func main() {
	cursos := Cursos{
		{Nome: "Go", CargaHoraria: 8},
		{Nome: "Java", CargaHoraria: 16},
		{Nome: "JavaScript", CargaHoraria: 20},
	}
	// Usa template.Must para simplificar o tratamento de erros
	// Carrega o template de um arquivo externo
	templateMust := template.Must((template.New("template.html")).ParseFiles("template.html"))
	// Executando o template
	err := templateMust.Execute(os.Stdout, cursos)
	if err != nil {
		panic(err)
	}
}
