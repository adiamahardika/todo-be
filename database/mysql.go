package database

import (
	"fmt"
	"log"
	"os"

	"svc-todo/entity"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func ConnectMysql() (*gorm.DB, error) {
	error_env := godotenv.Load()
	if error_env != nil {
		log.Fatal("Error loading .env file")
	}
	host := os.Getenv("MYSQL_HOST")
	user := os.Getenv("MYSQL_USER")
	port := os.Getenv("MYSQL_PORT")
	pass := os.Getenv("MYSQL_PASSWORD")
	db := os.Getenv("MYSQL_DBNAME")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", user, pass, host, port, db)
	return gorm.Open(mysql.Open(dsn), &gorm.Config{})

}

func AutoMigrateMysql(db *gorm.DB) {
	db.AutoMigrate(&entity.Activity{}, &entity.Todo{})

	db.Migrator().CreateIndex(&entity.Activity{}, "id")
	db.Migrator().CreateIndex(&entity.Todo{}, "id")
}
