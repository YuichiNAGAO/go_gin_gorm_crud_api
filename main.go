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

	r.POST("/posts", controllers.PostsCreate)
	// curl -X POST -H "Content-Type: application/json" -d '{"title":"タイトル","body":"本文"}' http://localhost:3000/posts

	r.GET("/posts", controllers.PostsIndex)

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
