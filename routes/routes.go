package routes

import (
	"net/http"
	"receitas_toto/controllers"
)

func Rotas() {
    http.HandleFunc("/", controllers.Index)
	http.HandleFunc("/login", controllers.PaginaLogin)
	http.HandleFunc("/receitas", controllers.PaginaReceitas)
}