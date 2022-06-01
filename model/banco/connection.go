package service

import (
	"database/sql"
	congif "ec2/model/config"
	"fmt"
	"log"

	_ "github.com/lib/pq" // import do drive do postgres, ele nao é usado nesse pacote por isso tem que definir ele como import explicito, usando o _ na frente do pacote
)

func ConectaDB() (*sql.DB, error) {
	congif.VariveisAm()
	conect := fmt.Sprintf("host=%s port=%d user=%s password=%v dbname=%s sslmode=disable ", congif.HOST, congif.DB_PORT, congif.DB_USER, congif.PASSWORD, congif.DB_NAME)

	db, err := sql.Open("postgres", conect)
	if err != nil {
		log.Fatal("erro ao conetar")
	}

	err = db.Ping()
	if err != nil {
		log.Fatal("erro no ping")
	}

	fmt.Printf("conectou-se ao banco de dados postgres in docker :) \n")

	return db, nil
}
