package main

import (
	"go_curso_3/routes"
	"net/http"
)

func main() {
	routes.CarregaRoutes()
	http.ListenAndServe(":8000", nil)
}
