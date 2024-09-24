package persistencia

import (
	"go_curso_3/persistenciaGenerica"
	"log"
)

type Produto struct {
	Nome, Descricao string
	Preco           float64
	Quantidade      int
	Id              int64
}

func FindAll() []Produto {

	db := persistenciaGenerica.Conectar()
	var listaProduto []Produto
	rows, err := db.Query("SELECT * FROM produto")
	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()
	defer rows.Close()

	for rows.Next() {
		var p Produto
		err := rows.Scan(&p.Id, &p.Nome, &p.Descricao, &p.Preco, &p.Quantidade)
		if err != nil {
			log.Fatal(err)
		}
		listaProduto = append(listaProduto, p)
	}
	return listaProduto
}

func Save(p Produto) {
	db := persistenciaGenerica.Conectar()
	defer db.Close()
	insere, err := db.Prepare("insert into produto (nome, descricao, preco, quantidade) values (?,?,?,?) ")
	if err != nil {
		log.Fatal(err)
	}
	insere.Exec(p.Nome, p.Descricao, p.Preco, p.Quantidade)
}

func Delete(codigo string) {
	db := persistenciaGenerica.Conectar()
	defer db.Close()
	delete, err := db.Prepare("delete from produto where id = ?")
	if err != nil {
		log.Fatal(err)
	}
	delete.Exec(codigo)
}
