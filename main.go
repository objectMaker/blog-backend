package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/joho/godotenv"
	"github.com/objectMaker/blog-backend/db"
	"github.com/objectMaker/blog-backend/router"
)

func main() {
	loadEnv()
	db.Connect()
	var isMigrate bool
	flag.BoolVar(&isMigrate, "migrate", false, "migrate database")
	flag.Parse()
	fmt.Println(isMigrate, "ismigrate")
	if isMigrate {
		err := db.Migration()
		if err != nil {
			log.Fatal("migration failed: ", err)
		} else {
			// 迁移成功后退出程序
			log.Println("migration success")
		}
		return
	}

	router.New()
}

func loadEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("env load failed: ", err)
	}
}
