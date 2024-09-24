package produto

import (
	"fmt"
	"go_curso_3/persistencia"
	"net/http"
	"strconv"
)

func FindAll() []persistencia.Produto {
	return persistencia.FindAll()
}

func Insert(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.FormValue("nome"))
	fmt.Println(r.FormValue("descricao"))
	if r.Method == "POST" {
		preco, _ := strconv.ParseFloat(r.FormValue("preco"), 64)
		qtd, _ := strconv.Atoi(r.FormValue("quantidade"))

		persistencia.Save(persistencia.Produto{
			Nome:       r.FormValue("nome"),
			Descricao:  r.FormValue("descricao"),
			Preco:      preco,
			Quantidade: qtd,
		})
	}
	http.Redirect(w, r, "/", 301)
}

func Delete(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		persistencia.Delete(r.URL.Query().Get("id"))
	}
	http.Redirect(w, r, "/", 301)
}
