package person

import (
	"database/sql"
	_ "github.com/lib/pq"
	"github.com/nathanpiina/golang_api_learning/database"
	"log"
)

var Db *sql.DB

func Main() {
	Db = database.DatabaseConnection()
	AddPeople("", "", "", "")
}

func AddPeople(nickname string, name string, birth string, stack string) {
	p := People{Nickname: nickname, Name: name, Birth: birth, Stack: stack}

	_, err := Db.Exec("INSERT INTO people (nickname, name, stack, birth) VALUES ($1, $2, $3, $4)", p.Nickname, p.Name, p.Stack, p.Birth)

	if err != nil {
		log.Fatal(err)
	}
}
