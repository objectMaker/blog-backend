package main

import (
	"log"

	"github.com/joho/godotenv"
	"github.com/objectMaker/blog-backend/db"
	"github.com/objectMaker/blog-backend/router"
)

func main() {
	loadEnv()
	db.Connect()
	router.New()
}

func loadEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("env load failed: ", err)
	}
}
