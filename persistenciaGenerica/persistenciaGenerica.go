package persistenciaGenerica

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

func Conectar() *sql.DB {
	db, err := sql.Open("mysql", "root:senha@tcp(127.0.0.1:3306)/goWeb")

	if err != nil {
		log.Fatal(err)
		return nil
	}
	fmt.Println("Conectado com sucesso ao MySQL!")
	return db
}
