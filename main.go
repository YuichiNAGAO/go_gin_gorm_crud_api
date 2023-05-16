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

	r.GET("/posts/:id", controllers.PostsShow)

	r.PUT("/posts/:id", controllers.PostsUpdate)
	// curl -X PUT -H "Content-Type: application/json" -d '{"title":"タイトルああああ","body":"本文だよ"}' http://localhost:3000/posts/1

	r.DELETE("/posts/:id", controllers.PostsDelete)
	// curl -X DELETE http://localhost:3000/posts/5

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
