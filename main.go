package main

import (
	"flag"
	"log"

	"github.com/joho/godotenv"
	"github.com/objectMaker/blog-backend/db"
	"github.com/objectMaker/blog-backend/router"
)

var isMigrate bool

func main() {
	loadEnv()
	db.Connect()
	db.Migrate()
	router.New()
}

func loadEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("env load failed: ", err)
	}
}

func parseFlag() {
	flag.BoolVar(&isMigrate, "migrate", false, "migrate database")
	flag.Parse()
}
