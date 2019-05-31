package user

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"

	"github.com/adifahmi/learn-gin/database/models"
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
	type RequestBody struct {
		Username string `form:"username" binding:"required,exists,alphanum,min=4,max=190"`
		Email    string `form:"email" binding:"required,exists,email,max=190"`
		Age      int    `form:"age" binding:"numeric"`
		Password string `form:"password" binding:"required"`
	}

	var body RequestBody
	if err := c.BindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "invalid request body",
		})
		fmt.Println(err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": "ok",
	})
}
