package routes

import (
	"go_curso_3/produto"
	"html/template"
	"net/http"
)

var temp = template.Must(template.ParseGlob("paginasHtml/*.html"))

func CarregaRoutes() {
	http.HandleFunc("/", index)
	http.HandleFunc("/new", new)
	http.HandleFunc("/insertProduto", InsertProduto)
	http.HandleFunc("/deletar", Deletar)
}
func index(w http.ResponseWriter, r *http.Request) {
	listaProduto := produto.FindAll()
	temp.ExecuteTemplate(w, "index", listaProduto)
}
func new(w http.ResponseWriter, r *http.Request) {
	temp.ExecuteTemplate(w, "novoProduto", nil)
}
func InsertProduto(w http.ResponseWriter, r *http.Request) {
	produto.Insert(w, r)
}
func Deletar(w http.ResponseWriter, r *http.Request) {
	produto.Delete(w, r)
}
