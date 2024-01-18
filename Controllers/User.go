package Controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/joshbedo/serverless-go/Models"
)

func ListAllUsers(c *gin.Context) {
	var users []Models.User
	err := Models.GetAllUsers(&users)
	if err != nil {
		// return message no users found
		c.JSON(404, gin.H{
			"message": "No users found",
		})
	} else {
		c.JSON(200, users)
	}
}

func GetUser(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "GetUser",
	})
}
