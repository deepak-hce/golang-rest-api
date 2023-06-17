package main

import (
	"rest-api/initializers"
	route "rest-api/routes"

	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

func main() {
	r := gin.Default()
	route.UserRoute(r)
	r.Run()
}
