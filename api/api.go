package api

import (
	"net/http"

	"github.com/adifahmi/learn-gin/api/user"
	"github.com/gin-gonic/gin"
)

func ping(c *gin.Context) {
	c.String(http.StatusOK, "pong")
}

func Routes(r *gin.Engine) {
	api := r.Group("/api")
	{
		api.GET("/ping", ping)
		user.Routes(api)
	}
}
