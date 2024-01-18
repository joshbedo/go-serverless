package Routers

import (
	"github.com/gin-gonic/gin"
	"github.com/joshbedo/serverless-go/Controllers"
)

func SetupRouters() *gin.Engine {
	r := gin.Default()

	v1 := r.Group("/v1")
	{
		v1.GET("/users", Controllers.ListAllUsers)
		v1.GET("/users/:id", Controllers.GetUser)
	}

	return r
	// r.GET("/ping", func(c *gin.Context) {
	// 	c.JSON(200, gin.H{
	// 		"message": "Hello, World!",
	// 	})
	// })

	// r.GET("/ping/:name", func(c *gin.Context) {
	// 	name := c.Param("name")

	// 	c.JSON(200, gin.H{
	// 		"message": fmt.Sprintf("Hello, %s!", name),
	// 	})
	// })

	// r.GET("/hello/:name", func(c *gin.Context) {
	// 	name := c.Param("name")

	// 	c.String(200, "Hello, %s!", name)
	// })
}
