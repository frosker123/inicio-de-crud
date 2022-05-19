package congif

import (
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
	SECRET   = ""
)

func VariveisAm() {
	var err error

	err = gotenv.Load("../.env")
	if err != nil {
		panic(err)
	}

	DB_PORT, err = strconv.Atoi(os.Getenv("DB_PORT"))
	if err != nil {
		panic(err)
	}

	PASSWORD, err = strconv.Atoi(os.Getenv("PASSWORD"))
	if err != nil {
		panic(err)
	}

	API_PORT, err = strconv.Atoi(os.Getenv("API_PORT"))
	if err != nil {
		panic(err)
	}

	HOST = os.Getenv("HOST")

	DB_NAME = os.Getenv("DB_NAME")

	DB_USER = os.Getenv("DB_USER")

	SECRET = os.Getenv("SECRET")

}
