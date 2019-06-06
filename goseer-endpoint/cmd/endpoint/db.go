package main

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type (
	tempModel struct {
		gorm.Model
		Temps string `json:"temps"`
	}
)

var db *gorm.DB

func initDB() {
	var err error
	db, err = gorm.Open("mysql", "")

	if err != nil {
		panic("failed to connect database")
	}

	db.AutoMigrate(&tempModel{})
}


