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

	// Criando um template
	template := template.New("CursoTemplate")
	// Parse do template
	template, _ = template.Parse("Curso: {{.Nome}} - Carga Horária: {{.CargaHoraria}} horas")
	// Executando o template
	// Ele preenche com base no struct passado
	err := template.Execute(os.Stdout, curso)
	if err != nil {
		panic(err)
	}
}
