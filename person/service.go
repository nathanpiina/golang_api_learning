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
}

func AddPeople(nickname string, name string, birth string, stack string) error {
	p := People{Nickname: nickname, Name: name, Birth: birth, Stack: stack}

	_, err := Db.Exec("INSERT INTO people (nickname, name, stack, birth) VALUES ($1, $2, $3, $4)", p.Nickname, p.Name, p.Stack, p.Birth)

	if err != nil {
		log.Fatal(err)
		return err
	}

	return nil
}

func SearchPeople(nickname string) (string, error) {
	row := Db.QueryRow("SELECT name, birth, stack FROM people WHERE nickname = ($1)", nickname)

	var name string
	var birth string
	var stack string

	err := row.Scan(&name, &birth, &stack)
	if err == sql.ErrNoRows {
		return "Nenhuma linha correspondente encontrada.", nil
	} else if err != nil {
		return "", err
	}

	return fmt.Sprintf("Nome: %s, Nascimento: %s, Stack: %s", name, birth, stack), nil
}

func CountRowsInTable() (string, error) {
	var rowCount int
	row := Db.QueryRow("SELECT COUNT(*) FROM people")

	err := row.Scan(&rowCount)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("Total de linhas na tabela: %d", rowCount), nil
}