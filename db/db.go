package db

import (
	"fmt"
	"log"

	"github.com/objectMaker/blog-backend/models"
	"github.com/objectMaker/blog-backend/tools"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	var err error
	dbEnvStruct := struct {
		Host     string `envField:"host"`
		User     string `envField:"user"`
		Password string `envField:"password"`
		Dbname   string `envField:"dbname"`
		Port     string `envField:"dbport"`
	}{}
	tools.LoadEnvByStruct(&dbEnvStruct)
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Shanghai", dbEnvStruct.Host, dbEnvStruct.User, dbEnvStruct.Password, dbEnvStruct.Dbname, dbEnvStruct.Port)
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to connect database: %v", err)
	}
	log.Println("database connected")
}

func Migrate() {
	err := DB.AutoMigrate(&models.User{})
	if err != nil {
		log.Fatal("migration failed: ", err)
	} else {
		// 迁移成功后退出程序
		log.Println("migration success")
	}
}
