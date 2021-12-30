package authcontroller

import (
	"log"
	"net/http"
	"togolist_gin/models"
	"togolist_gin/services/auth"

	"github.com/gin-gonic/gin"
)

func Signup(c *gin.Context){
	newUser := models.UserModel{}
	if err := c.ShouldBindJSON(&newUser); err != nil {
		log.Fatal(err)
		c.Status(http.StatusBadRequest)
	} 
	
	if checkUsername := authservice.Checkusername(newUser.Username); !checkUsername{
		c.JSON(http.StatusOK, gin.H{
			"message": "This username is already used.",
			"username": newUser.Username,
		})
	}else{
		authservice.CreateUser(&newUser)
		c.JSON(http.StatusOK, gin.H{
			"message": "Created.",
			"ID": newUser.ID,
			"username": newUser.Username,
		})
	}
}

func Signin(c *gin.Context){
	signinUser := models.SigninModel{}
	if err := c.ShouldBindJSON(&signinUser); err != nil{
		log.Fatal(err)
		c.Status(http.StatusBadRequest)
	}
	user, signinResult := authservice.Signin(signinUser)
	if signinResult {
		accessToken, _ := authservice.GenerateToken(user)
		c.JSON(http.StatusOK, gin.H{
			"message": "Successful.",
			"AccessToken": accessToken,
		})
	}else{
		c.JSON(http.StatusOK, gin.H{
			"message": "Username or Password is invalid.",
		})
	}
}

func UserInfo(c *gin.Context){
	user_id, _ := c.Get("user_id")
	
	user, err := authservice.UserInfo(user_id.(float64))
	if err != nil{
		c.Status(http.StatusUnauthorized)
	}
	c.JSON(http.StatusOK, gin.H{
		"id": user.ID,
		"name": user.Name,
		"username": user.Username,
	})
}