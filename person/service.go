package person

import (
	"database/sql"
	"fmt"
	"log"
	_ "github.com/lib/pq"
	"github.com/nathanpiina/golang_api_learning/database"
)

var Db *sql.DB

func Main() {
	Db = database.DatabaseConnection()
	CountRowsInTable()
}

func AddPeople(nickname string, name string, birth string, stack string) {
	p := People{Nickname: nickname, Name: name, Birth: birth, Stack: stack}

	_, err := Db.Exec("INSERT INTO people (nickname, name, stack, birth) VALUES ($1, $2, $3, $4)", p.Nickname, p.Name, p.Stack, p.Birth)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Usuario adicionado com sucesso")
}

func SearchPeople(nickname string) {

	row := Db.QueryRow("SELECT name, birth, stack FROM people WHERE nickname = ($1)", nickname)

	var name string
	var birth string
	var stack string

	err := row.Scan(&name, &birth, &stack)
	if err == sql.ErrNoRows {
		fmt.Println("Nenhuma linha correspondente encontrada.")
	} else if err != nil {
		fmt.Println("Erro ao escanear resultados:", err)
	} else {
		fmt.Printf("Nome: %s, Nascimento: %s, Stack: %s\n", name, birth, stack)
	}
}

func CountRowsInTable() {
	var rowCount int
	row := Db.QueryRow("SELECT COUNT(*) FROM people")

	err := row.Scan(&rowCount)
	if err != nil {
		fmt.Println("Erro ao escanear resultados:", err)
	} else {
		fmt.Printf("Total de linhas na tabela: %d\n", rowCount)
	}
}
