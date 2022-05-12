package service

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

func ConectaDB() (*sql.DB, error) {
	var (
		HOST     = "localhost"
		PORT     = 5432
		USER     = "postgres"
		PASSWORD = 12345
		DB_NAME  = "postgresql"
	)
	conect := fmt.Sprintf("host=%s port=%d user=%s password=%v dbname=%s sslmode=disable ", HOST, PORT, USER, PASSWORD, DB_NAME)

	db, err := sql.Open("postgres", conect)
	if err != nil {
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	fmt.Printf("consegui conectar no banco de dados :)")

	return db, nil
}
