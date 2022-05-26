package utils

import (
	"html/template"
	"net/http"
)

var templates *template.Template

//carrega as telas web que foram criandas na views
func CarregaTemplate() {
	templates = template.Must(templates.ParseGlob("views/*.html"))
}

//executa as telas web que carregadas no method carregaTemplate
func ExecTemplate(w http.ResponseWriter, templete string, dados interface{}) {
	templates.ExecuteTemplate(w, templete, dados)
}
