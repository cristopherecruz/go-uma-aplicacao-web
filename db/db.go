package db

import "database/sql"
import _ "github.com/lib/pq"

func ConectarBancoDados() *sql.DB {

	connStr := "user=postgres dbname=loja password=1234 host=localhost sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		panic(err.Error())
	}

	return db
}
