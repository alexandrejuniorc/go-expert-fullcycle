package main

import (
	"html/template"
	"net/http"
)

type Curso struct {
	Nome         string
	CargaHoraria int
}

type Cursos []Curso

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		cursos := Cursos{
			{Nome: "Go", CargaHoraria: 8},
			{Nome: "Java", CargaHoraria: 16},
			{Nome: "JavaScript", CargaHoraria: 20},
		}

		templateMust := template.Must((template.New("template.html")).ParseFiles("template.html"))
		err := templateMust.Execute(w, cursos)
		if err != nil {
			panic(err)
		}
	})

	http.ListenAndServe(":8282", nil)
}
