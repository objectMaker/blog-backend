package main

import (
	"fmt"
	"log"

	"github.com/joho/godotenv"
	"github.com/objectMaker/blog-backend/db"
)

func init() {
	loadEnv()
	db.Connect()
}

func main() {
	fmt.Println("hello blog")
}

func loadEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("env load failed: ", err)
	}
}
