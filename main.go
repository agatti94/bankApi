package main

import (
	"html/template"
	"net/http"
)

type Account struct {
	Id         int
	Name       string
	Cpf        string
	Secret     string
	Balance    float64
	Created_at string
}

// encapsulando templates, passando endereço dos templates e carregando
var temp = template.Must(template.ParseGlob("templates/*.html"))

func main() {
	http.HandleFunc("/", index)
	http.ListenAndServe(":8000", nil)
}

func index(w http.ResponseWriter, r *http.Request) {
	accounts := []Account{
		{23, "Arthur", "349.932.308-77", "logan1290", 1424.23, "09/09/2021"},
		{53, "Pedro", "097.194.508-00", "rosebud", 1233.23, "08/09/2021"},
	}
	temp.ExecuteTemplate(w, "Index", accounts)
}
