package controllers

import (
	"fmt"
	"github.com/ganggas95/trawanganserver/app/models"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"log"
)

type Database struct {
}

const (
	DB_USER = "postgres"
	DB_PASS = "K3yf4t0n"
	DB_NAME = "db_prog"
)

func CheckError(err error) {
	if err != nil {
		log.Println(err)
	}
}

func (d Database) initDb() *gorm.DB {
	dbinfo := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable", DB_USER, DB_PASS, DB_NAME)
	db, err := gorm.Open("postgres", dbinfo)
	if err != nil {
		CheckError(err)
	}
	d.createTable(db)
	return db
}
func (d Database) createTable(db *gorm.DB) {
	db.SingularTable(true)
	db.AutoMigrate(&models.AgentTravel{},
		&models.User{},
		&models.UserToken{},
		&models.Person{},
		&models.AgentService{},
		&models.UserFoto{},
		&models.FotoService{},
		&models.AddOnService{},
	)

}
