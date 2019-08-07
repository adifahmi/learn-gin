package user

import (
	"fmt"
	"net/http"
	// "reflect"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"

	"github.com/adifahmi/learn-gin/database/models"
	"github.com/adifahmi/learn-gin/lib"
	"github.com/mitchellh/mapstructure"
)

type User = models.User

func listUsers(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	var users []User

	err := db.Select("id, username, email, age").Find(&users).Error
	if err != nil {
		c.AbortWithStatus(500)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": users,
	})
}

func serializeAndGenerateToken(user User) (map[string]interface{}, string) {
	serializedUser := user.Serialize()
	token, tokenErr := lib.GenerateToken(serializedUser, 7)
	if tokenErr != nil {
		panic("Err generate token")
	}

	return serializedUser, token
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

	serializedUser, token := serializeAndGenerateToken(user)

	c.JSON(200, lib.JSON{
		"user":  serializedUser,
		"token": token,
	})
}

func login(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	type RequestBody struct {
		Username string `json:"username" binding:"required"`
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

	// Check existence
	var currentUser User
	if db.Where("username = ?", body.Username).First(&currentUser).RecordNotFound() {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "username doesn't exists",
		})
		return
	}

	if lib.CheckHash(body.Password, currentUser.Password) == false {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "invalid password",
		})
		return
	}

	serializedUser, token := serializeAndGenerateToken(currentUser)

	c.JSON(200, lib.JSON{
		"user":  serializedUser,
		"token": token,
	})
}

func check(c *gin.Context) {
	token := c.Request.Header.Get("token")

	if token == "" {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "empty token",
		})
		return
	}

	claims, err := lib.ParseToken(token)

	if err != nil {
		fmt.Println("err parse token", err)
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "invalid token",
		})
		return
	}

	var currentUser User
	mapstructure.Decode(claims["user"], &currentUser)

	c.JSON(200, lib.JSON{
		"token": token,
		"user":  currentUser.Serialize(),
	})
}
