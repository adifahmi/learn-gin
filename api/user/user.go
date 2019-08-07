package user

import "github.com/gin-gonic/gin"

// Routes list all avail route
func Routes(r *gin.RouterGroup) {
	user := r.Group("/user")
	{
		user.GET("/", listUsers)
		user.POST("/register", register)
		user.POST("/login", login)
		user.GET("/check", check)
	}
}
