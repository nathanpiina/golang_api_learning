package person

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"github.com/nathanpiina/golang_api_learning/database"
	"log"
)

var Db *sql.DB

func Main() {
	Db = database.DatabaseConnection()
	SearchPeople("nickname")
}

func AddPeople(nickname string, name string, birth string, stack string) {
	p := People{Nickname: nickname, Name: name, Birth: birth, Stack: stack}

	_, err := Db.Exec("INSERT INTO people (nickname, name, stack, birth) VALUES ($1, $2, $3, $4)", p.Nickname, p.Name, p.Stack, p.Birth)

	if err != nil {
		log.Fatal(err)
	}
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
