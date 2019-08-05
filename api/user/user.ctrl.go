package user

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"

	"github.com/adifahmi/learn-gin/database/models"
	"github.com/adifahmi/learn-gin/lib"
)

type User = models.User

func listUsers(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	var users []User

	err := db.Select("id, username, email, age").Find(&users).Error
	if err != nil {
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"data": users,
	})
}

func register(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	type RequestBody struct {
		Username string `json:"username" binding:"required,exists,alphanum,min=4,max=190"`
		Email    string `json:"email" binding:"required,exists,email,max=190"`
		Age      int    `json:"age"`
		Password string `json:"password" binding:"required"`
	}

	var body RequestBody
	if err := c.BindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "invalid request body",
		})
		fmt.Println(err)
		return
	}

	// Check duplication
	var exists User
	if db.Where("username = ?", body.Username).Or("email = ?", body.Email).First(&exists).RecordNotFound() == false {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "username or email already exists",
		})
		return
	}

	hash, hashErr := lib.Hash(body.Password)
	if hashErr != nil {
		fmt.Println("hashErr", hashErr)
		c.AbortWithStatus(500)
		return
	}

	fmt.Println(hash)

	// create user
	user := User{
		Username: body.Username,
		Email:    body.Email,
		Age:      body.Age,
		Password: hash,
	}

	db.NewRecord(user)
	db.Create(&user)

	serialized := user.Serialize()
	token, tokenErr := lib.GenerateToken(serialized, 7)
	if tokenErr != nil {
		fmt.Println("tokenErr", tokenErr)
		c.AbortWithStatus(500)
		return
	}

	c.SetCookie("token", token, 60*60*24*7, "/", "", false, true)

	c.JSON(200, lib.JSON{
		"user":  user.Serialize(),
		"token": token,
	})
}
