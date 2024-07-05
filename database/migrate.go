package database

import (
	"errors"
	"log"
	"webserver/model"
)

func Migrate() error {
	if GetDB() == nil {
		log.Fatal("could not connect to database")
		return errors.New("db is nil")
	}
	err := db.Table("work_category").AutoMigrate(&model.WorkCategory{})
	if err != nil {
		log.Fatal("could not migrate work_category")
		return err
	}
	log.Println("migration done successfully")
	return nil
}
