package controllers

import (
	"html/template"
	"net/http"
)

// Carrega os templates
var templates = template.Must(template.ParseGlob("templates/*.html")) 

// Exibe a tela index.html
func Index(w http.ResponseWriter, r *http.Request) {
    templates.ExecuteTemplate(w, "Index", nil)
}

// Carrega a página login.html
func PaginaLogin(w http.ResponseWriter, r *http.Request) {
	templates.ExecuteTemplate(w, "Login", nil)
}

// Carrega a página de receitas
func PaginaReceitas(w http.ResponseWriter, r *http.Request) {
	templates.ExecuteTemplate(w, "Receitas", nil)
}