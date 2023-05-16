package main

import (
	"github.com/YuichiNAGAO/go_gin_gorm_crud_api/controllers"
	"github.com/YuichiNAGAO/go_gin_gorm_crud_api/initializers"
	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectDB()
}

func main() {
	r := gin.Default()
	// r.GET("/ping", func(c *gin.Context) {
	// 	c.JSON(http.StatusOK, gin.H{
	// 		"message": "pong",
	// 	})
	// })

	r.POST("/posts", controllers.PostsCreate)

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
