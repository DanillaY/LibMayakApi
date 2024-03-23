package main

import (
	"fmt"
	"os"

	server "github.com/DanillaY/BookApi/cmd"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load("./db.env")
	if err != nil {
		fmt.Println("Error while getting env data")
	}

	config := server.Config{

		HOST:     os.Getenv("HOST"),
		DB_PORT:  os.Getenv("DB_PORT"),
		API_PORT: os.Getenv("API_PORT"),
		PASSWORD: os.Getenv("PASSWORD"),
		DB:       os.Getenv("DB_NAME"),
		USER:     os.Getenv("USER"),
		SSLMODE:  os.Getenv("SSLMODE_TYPE"),
	}

	db, err := server.NewPostgresConnection(&config)
	repo := server.Repository{Db: db, Config: &config}
	repo.InitAPIRoutes()
}