package main

import (
	"html/template"
	"os"
)

type Curso struct {
	Nome         string
	CargaHoraria int
}

func main() {
	curso := Curso{
		Nome:         "Go Expert",
		CargaHoraria: 8,
	}
	// Usa template.Must para simplificar o tratamento de erros
	templateMust := template.Must((template.New("CursoTemplate")).Parse("Curso: {{.Nome}} - Carga Horária: {{.CargaHoraria}} horas"))
	// Executando o template
	err := templateMust.Execute(os.Stdout, curso)
	if err != nil {
		panic(err)
	}
}
