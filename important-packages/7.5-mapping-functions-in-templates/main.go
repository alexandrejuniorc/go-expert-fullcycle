package main

import (
	"html/template"
	"os"
	"strings"
)

type Curso struct {
	Nome         string
	CargaHoraria int
}

type Cursos []Curso

func ToUpper(s string) string {
	return strings.ToUpper(s)
}

func main() {
	templates := []string{"header.html", "content.html", "footer.html"}

	cursos := Cursos{
		{Nome: "Go", CargaHoraria: 8},
		{Nome: "Java", CargaHoraria: 16},
		{Nome: "JavaScript", CargaHoraria: 20},
	}

	// criando um novo template e mapeando a função ToUpper
	t := template.New("content.html")
	// mapeando a função ToUpper
	t.Funcs(template.FuncMap{"ToUpper": ToUpper})
	// parse dos arquivos
	t = template.Must(t.ParseFiles(templates...))

	err := t.Execute(os.Stdout, cursos)
	if err != nil {
		panic(err)
	}

}
