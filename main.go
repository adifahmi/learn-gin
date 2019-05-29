package main

import (
	"os"

	"github.com/adifahmi/learn-gin/api"
	"github.com/adifahmi/learn-gin/database"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}

	// Init the DB
	db, _ := database.DBInit()

	app := gin.Default()
	app.Use(database.Inject(db))
	api.Routes(app)
	app.Run(":" + os.Getenv("APP_PORT"))
}
