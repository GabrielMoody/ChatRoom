package mysql

import (
	"fmt"
	"github.com/GabrielMoody/chat-app/server/internal/helper"
	"github.com/gofiber/fiber/v2/log"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func NewConnection() *gorm.DB {
	v := helper.LoadEnv()

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", v.GetString("MYSQL_USERNAME"), v.GetString("MYSQL_PASSWORD"), v.GetString("MYSQL_HOST"), v.GetString("MYSQL_PORT"), v.GetString("MYSQL_DATABASE"))

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatalf("failed to connect to mysql: %v", err)
	}

	err = db.AutoMigrate(&User{})

	if err != nil {
		log.Fatalf("failed to auto migrate users: %v", err)
	}

	return db
}
