package database

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

func DatabaseConnection() *sql.DB {
	host := "localhost"
	port := 5432
	user := "postgres"
	password := "n82161743"
	dbname := "postgres"

	connectionString := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	Db, err := sql.Open("postgres", connectionString)
	if err != nil {
		panic(err)
	}

	err = Db.Ping()
	if err != nil {
		panic(err)
	}

	fmt.Println("Conexão com o banco de dados PostgreSQL estabelecida com sucesso!")

	return Db
}
