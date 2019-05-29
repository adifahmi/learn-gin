package user

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"

	"github.com/adifahmi/learn-gin/database/models"
)

type User = models.User

func listUsers(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	var users []User

	err := db.Find(&users).Error
	if err != nil {
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"data": users,
	})
}
