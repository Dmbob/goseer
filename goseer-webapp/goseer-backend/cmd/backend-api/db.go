package main

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

/*
* This function will initialize a database connection and then return it.
*/
func initDB(config *Config) *gorm.DB {
	DBUSER := config.DBUser
	DBPASS := config.DBPass
	DBADDRESS := config.DBAddress
	DBPORT := config.DBPort
	DBNAME := config.DBName

	db, err := gorm.Open("mysql", DBUSER+":"+DBPASS+"@tcp("+DBADDRESS+":"+DBPORT+")/"+DBNAME+"?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}

	migrateStructs(db)

	return db
}

func migrateStructs(db *gorm.DB) {
	db.AutoMigrate(&User{})
}
