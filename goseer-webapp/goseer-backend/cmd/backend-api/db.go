package main

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

/*
* This function will initialize a database connection and then return it.
*/
func initDB() *gorm.DB {
	DBUSER := ""
	DBPASS := ""
	DBADDRESS := ""
	DBPORT := ""
	DBNAME := "goseer"

	db, err := gorm.Open("mysql", DBUSER+":"+DBPASS+"@tcp("+DBADDRESS+":"+DBPORT+")/"+DBNAME+"?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&User{})

	return db
}
