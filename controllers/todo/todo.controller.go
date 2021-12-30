package todocontroller

import (
	"log"
	"net/http"
	"togolist_gin/models"
	"togolist_gin/services/todo"
	"strconv"

	"github.com/gin-gonic/gin"
)

func getIdFromToken(c *gin.Context) int{
	user_id, _ := c.Get("user_id")
	return int(user_id.(float64))
}

func InputTodo(c *gin.Context){
	id := getIdFromToken(c)
	newTodo := models.TodoModel{}
	newTodo.UserRefer = id
	if err := c.ShouldBindJSON(&newTodo); err != nil {
		log.Fatal(err)
		c.Status(http.StatusBadRequest)
	}
	todoservice.InputTodo(&newTodo)
	c.JSON(http.StatusOK, gin.H{
		"todo": newTodo,
	})
}

func UserTodo(c *gin.Context){
	id := getIdFromToken(c)
	userTodo, err := todoservice.GetUserTodo(id)
	if err != nil{
		c.Status(http.StatusInternalServerError)
	}else{
		c.JSON(http.StatusOK, gin.H{
			"todo": userTodo,
		})
	}
}

func DeleteTodo(c *gin.Context){
	userId := getIdFromToken(c)
	targetTodoId, err := strconv.Atoi(c.Param("id"))
	if err != nil{
		c.Status(http.StatusBadRequest)
	}
	message, err := todoservice.DeleteUserTodo(userId, targetTodoId)
	if err != nil{
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": message,
		})
	}else{
		c.JSON(http.StatusOK, gin.H{
			"message": message,
		})
	}
}

func UpdateTodo(c *gin.Context){
	userId := getIdFromToken(c)
	updateTodo := models.TodoModel{}
	c.ShouldBindJSON(&updateTodo)
	message, err := todoservice.UpdateUserTodo(updateTodo, userId)
	if err != nil{
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": message,
		})
	}else{
		c.JSON(http.StatusOK, gin.H{
			"message": message,
		})
	}
}