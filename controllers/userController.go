package controller

import (
	"fmt"
	"rest-api/initializers"
	"rest-api/models"
	"time"

	"github.com/gin-gonic/gin"
)

func CreateUser(c *gin.Context) {

	var body struct {
		FirstName   string
		LastName    string
		DateOfBirth string
		Photo       string
	}

	c.Bind(&body)

	fmt.Println(body.DateOfBirth)

	var date time.Time
	var error error
	dateString := body.DateOfBirth
	date, error = time.Parse("2006-01-02", dateString)
	if error != nil {
		c.JSON(400, gin.H{
			"message": "Error while creating user",
			"result":  error,
		})
		return
	}

	user := models.User{FirstName: body.FirstName, LastName: body.LastName, Photo: body.Photo, DateOfBirth: date}

	result := initializers.DB.Create(&user)

	if result.Error != nil {
		c.Status(400)
		return
	}

	c.JSON(200, gin.H{
		"message": "User created successfully",
		"result":  user,
	})
}

func FetchUsers(c *gin.Context) {

	var users []models.User

	result := initializers.DB.Limit(100).Find(&users)

	if result.Error != nil {
		c.Status(400)
		return
	}

	c.JSON(200, gin.H{
		"message": "User fetched successfully",
		"result":  users,
	})
}
