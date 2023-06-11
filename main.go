package main

import (
	controller "rest-api/controllers"
	"rest-api/initializers"

	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

func main() {
	r := gin.Default()
	r.POST("/user", controller.CreateUser)
	r.Run()
}
