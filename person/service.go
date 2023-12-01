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
        log.Printf("erro ao adicionar pessoa: %v", err)
        return fmt.Errorf("erro ao adicionar pessoa: %v", err)
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
        return "", fmt.Errorf("nenhuma pessoa encontrada com o nickname %s", nickname)
    } else if err != nil {
        return "", fmt.Errorf("erro ao buscar pessoa: %v", err)
    }

    return fmt.Sprintf("nome: %s, nascimento: %s, stack: %s", name, birth, stack), nil
}

func CountRowsInTable() (string, error) {
    var rowCount int
    row := Db.QueryRow("SELECT COUNT(*) FROM people")

    err := row.Scan(&rowCount)
    if err != nil {
        return "", fmt.Errorf("erro ao contar linhas na tabela: %v", err)
    }

    return fmt.Sprintf("total de linhas na tabela: %d", rowCount), nil
}