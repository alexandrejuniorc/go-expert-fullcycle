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
	templates := []string{"header.html", "content.html", "footer.html"}

	cursos := Cursos{
		{Nome: "Go", CargaHoraria: 8},
		{Nome: "Java", CargaHoraria: 16},
		{Nome: "JavaScript", CargaHoraria: 20},
	}

	templateMust := template.Must((template.New("content.html")).ParseFiles(templates...))
	err := templateMust.Execute(os.Stdout, cursos)
	if err != nil {
		panic(err)
	}

}
