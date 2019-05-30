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
	db, err = gorm.Open("mysql", "root:eh9Y2WGY89fyZ@tcp(dmbob.guru:2222)/goseer?charset=utf8&parseTime=True&loc=Local")

	if err != nil {
		panic("failed to connect database")
	}

	db.AutoMigrate(&tempModel{})
}


