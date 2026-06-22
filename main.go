package main

import (
	"net/http"
	"receitas_toto/routes"
)

func main() {

	// Carrega arquivos estáticos (fotos, css, etc.)
	fs := http.FileServer(http.Dir("static/"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	routes.Rotas()
    http.ListenAndServe(":8000", nil)
}