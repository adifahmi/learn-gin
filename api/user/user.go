package user

import "github.com/gin-gonic/gin"

func Routes(r *gin.RouterGroup) {
	user := r.Group("/user")
	{
		user.GET("/", listUsers)
	}
}
