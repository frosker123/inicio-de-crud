package congif

import (
	"log"
	"os"
	"strconv"

	"github.com/subosito/gotenv"
)

var (
	DB_PORT  = 0
	API_PORT = 0
	PASSWORD = 0
	HOST     = ""
	DB_USER  = ""
	DB_NAME  = ""
	SECRET   []byte
)

func VariveisAm() {
	var err error

	err = gotenv.Load("../.env")
	if err != nil {
		log.Fatal("erro variavel")
	}

	DB_PORT, err = strconv.Atoi(os.Getenv("DB_PORT"))
	if err != nil {
		log.Fatal("erro variavel")
	}

	PASSWORD, err = strconv.Atoi(os.Getenv("PASSWORD"))
	if err != nil {
		log.Fatal("erro variavel")
	}

	API_PORT, err = strconv.Atoi(os.Getenv("API_PORT"))
	if err != nil {
		log.Fatal("erro variavel")
	}

	HOST = os.Getenv("HOST")

	DB_NAME = os.Getenv("DB_NAME")

	DB_USER = os.Getenv("DB_USER")

	SECRET = []byte(os.Getenv("SECRET"))

}
