package route

import (
	controller "rest-api/controllers"

	"github.com/gin-gonic/gin"
)

func UserRoute(r *gin.Engine) {
	r.POST("/user", controller.CreateUser)
	r.GET("/user", controller.FetchUsers)

}
