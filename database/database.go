package database

import (
	"fmt"
	"os"

	"github.com/adifahmi/learn-gin/database/models"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql" // mysql
)

// DBInit will init the DB
func DBInit() (*gorm.DB, error) {
	dbConfig := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PWD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"),
	)
	db, err := gorm.Open("mysql", dbConfig)
	db.LogMode(true)
	if err != nil {
		panic(err)
	}
	fmt.Println("Connected to DB")
	models.Migrate(db)
	return db, err
}

// Inject injects database to gin context
func Inject(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set("db", db)
		c.Next()
	}
}
