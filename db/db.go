package db

import "database/sql"

func conectaComBancoDeDados() *sql.DB {
	conexao := "user=postgres dbname=alura_loja password=4S06iV0w host=localhost sslmode=disable"
	db, err := sql.Open("postgres", conexao)
	if err != nil {
		panic(err.Error())
	}
	return db
}
