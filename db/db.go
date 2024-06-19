package db

import (
	"fmt"

	"github.com/objectMaker/blog-backend/tools"
)

//   dsn := "host=localhost user=gorm password=gorm dbname=gorm port=9920 sslmode=disable TimeZone=Asia/Shanghai"
//   db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

func Connect() {
	dbEnvStruct := struct {
		Host     string `envField:"host"`
		User     string `envField:"user"`
		Password string `envField:"password"`
		Dbname   string `envField:"dbname"`
		Port     string `envField:"port"`
	}{}

	tools.LoadEnvByStruct(&dbEnvStruct)
	fmt.Println(dbEnvStruct)
}
