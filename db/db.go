package db

import (
	"database/sql"

	_ "github.com/lib/pq"
)

func ConnectDB() *sql.DB {
	// connection := "user dbname password host sslmode"
	connection := "user=postgres dbname=loja password=root host=172.17.0.2 sslmode=disable"
	db, err := sql.Open("postgres", connection)
	if err != nil {
		panic(err.Error())
	}

	return db
}
