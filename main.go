package main

import (
	"togolist_gin/config"
	"togolist_gin/routes"
	"github.com/gin-gonic/gin"
)

func main(){
	r := gin.Default()
	routes.Route(r)
	config.ConnectDb()
	r.Run()
}