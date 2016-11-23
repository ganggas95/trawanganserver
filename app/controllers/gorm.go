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
	DB_HOST = "ec2-107-22-251-225.compute-1.amazonaws.com"
	DB_USER = "vcpyraxvhxkcxc"
	DB_PASS = "gHkAPj8weKd4a8_KGH_y1-bOYh"
	DB_NAME = "daalefbcj1ekfm"
)

func CheckError(err error) {
	if err != nil {
		log.Println(err)
	}
}

func (d Database) initDb() *gorm.DB {
	dbinfo := fmt.Sprintf("host=%s user=%s password=%s dbname=%s sslmode=verify-full", DB_HOST, DB_USER, DB_PASS, DB_NAME)
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
