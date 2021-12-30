package routes

import (
	"togolist_gin/controllers/auth"
	"togolist_gin/middleware"
	"togolist_gin/controllers/todo"

	"github.com/gin-gonic/gin"
)

func Route(r *gin.Engine){
	auth := r.Group("/auth")
	{
		auth.POST("/signup", authcontroller.Signup)
		auth.POST("/signin", authcontroller.Signin)
		auth.GET("/info", middleware.AuthorizeJWT(), authcontroller.UserInfo)
	}
	todo := r.Group("/todo")
	todo.Use(middleware.AuthorizeJWT())
	{
		todo.POST("/create", todocontroller.InputTodo)
		todo.GET("/usertodo", todocontroller.UserTodo)
		todo.DELETE("/deletetodo/:id", todocontroller.DeleteTodo)
		todo.PATCH("/updatetodo", todocontroller.UpdateTodo) //ยังไม่เทสนะ
	}
}