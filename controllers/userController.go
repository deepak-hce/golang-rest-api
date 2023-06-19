package controller

import (
	"encoding/json"
	"fmt"
	"rest-api/initializers"
	"rest-api/interfaces"
	"rest-api/models"
	"strconv"
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

	var itemsPerPage int
	var pageNumber int

	itemsPerPageQuery := c.Query("itemsPerPage")
	pageNumberQuery := c.Query("pageNumber")

	if itemsPerPageQuery == "" {
		itemsPerPage = 10
	} else {
		s, err := strconv.Atoi(itemsPerPageQuery)
		if err != nil {
			c.JSON(400, gin.H{
				"error": err,
			})
			fmt.Println(err)
			return
		}
		itemsPerPage = s
	}

	if pageNumberQuery == "" {
		pageNumber = 1
	} else {
		s, err := strconv.Atoi(pageNumberQuery)
		if err != nil {
			c.JSON(400, gin.H{
				"error": err,
			})
			fmt.Println(err)

			return
		}

		pageNumber = s
	}

	var users []models.User

	result := initializers.DB.Order("created_at DESC").Limit(itemsPerPage).Offset((pageNumber - 1) * itemsPerPage).Find(&users)

	if result.Error != nil {
		c.Status(400)
		return
	}

	encodedUsers, err := json.Marshal(&users)

	if err != nil {
		fmt.Println(err)
	}

	var marshalUsers []interfaces.User
	err = json.Unmarshal(encodedUsers, &marshalUsers)

	if err != nil {
		fmt.Println(err)
	}

	c.JSON(200, gin.H{
		"message": "User fetched successfully",
		"result":  marshalUsers,
	})
}
