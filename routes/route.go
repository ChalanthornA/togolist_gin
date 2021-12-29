package routes

import (
	"togolist_gin/controllers/auth"
	"togolist_gin/middleware"

	"github.com/gin-gonic/gin"
)

func Route(r *gin.Engine){
	auth := r.Group("/auth")
	{
		auth.POST("/signup", authcontroller.Signup)
		auth.POST("/signin", authcontroller.Signin)
	}
	user := r.Group("/user")
	user.Use(middleware.AuthorizeJWT())
	{
		user.GET("/getuser", authcontroller.Test)
	}
}