package main

import (
	"fmt"
	"log"

	"github.com/joho/godotenv"
)

func main() {
	fmt.Println("hello blog")
}

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("env load failed: ", err)
	}
}
