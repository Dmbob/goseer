package main

import (
	"golang.org/x/crypto/bcrypt"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

type (
	// User structure used for database interactions.
	User struct {
		gorm.Model
		Username     string `json:"username"`
		PasswordHash string	`json:"pass"`
	}
)

/*
* This function will create a user in the database if they do not exist already.
*/
func createUser(db *gorm.DB) gin.HandlerFunc {
	fn := func(c *gin.Context) {
		var user User
		c.BindJSON(&user)

		var existingUsers []User
	
		db.Where("username = ?", user.Username).Find(&existingUsers)

		if len(existingUsers) == 0 {
			hash, err := bcrypt.GenerateFromPassword([]byte(user.PasswordHash), bcrypt.MinCost)
			if err != nil {
				panic(err)
			}

			user.PasswordHash = string(hash)

			db.Create(&user)
			c.JSON(201, gin.H{"message": "User created successfully"})
		}else {
			c.JSON(202, gin.H{"message": "User already exists"})
		}
	}

	return gin.HandlerFunc(fn)
}

/*
* This function will verify a user's password and return a login token if it is successful.
*/
func verifyUser(db *gorm.DB) gin.HandlerFunc {
	fn := func(c *gin.Context) {
		var user User
		c.BindJSON(&user)

		var foundUser User

		db.Where("username = ?", user.Username).First(&foundUser)

		err := bcrypt.CompareHashAndPassword([]byte(foundUser.PasswordHash), []byte(user.PasswordHash))

		if err != nil {
			c.JSON(401, gin.H{"message": "Unauthorized"})
			return
		}

		c.JSON(202, gin.H{"token": "12345"}) // Return authentication token for the user.
		return
	}

	return gin.HandlerFunc(fn)
}