package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

var (
	Port  string
	DBUrl string
)

func init() {
	err := godotenv.Load("./pkg/config/.env")
	if err != nil {
		log.Fatal(err.Error())
	}
	db, isExist := os.LookupEnv("DB_URL")
	if !isExist {
		log.Fatal("Could not find DB_URL in environment variable")
	}
	DBUrl = db
	port, isExist := os.LookupEnv("PORT")
	if !isExist {
		log.Fatal("Could not find PORT in environment variable")
	}
	Port = port
}
