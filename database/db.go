package database

import (
	"github.com/RenanLourenco/go-gin.git/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	DB *gorm.DB
	err error
)


func ConectarDatabase(){
	dsn := "host=localhost user=admin password=root dbname=go-gin-rest port=5432 sslmode=disable TimeZone=America/Sao_Paulo"
	DB, err = gorm.Open(postgres.Open(dsn))
	if err != nil {
		panic(err.Error())
	}
	DB.AutoMigrate(&models.Aluno{})
}