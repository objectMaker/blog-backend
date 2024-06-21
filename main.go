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

	parseFlag()

	if isMigrate {
		migrate()
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

func parseFlag() {
	flag.BoolVar(&isMigrate, "migrate", false, "migrate database")
	flag.Parse()
}

func migrate() {
	err := db.Migration()
	if err != nil {
		log.Fatal("migration failed: ", err)
	} else {
		// 迁移成功后退出程序
		log.Println("migration success")
	}
}
