package controllers

import (
	"github.com/YuichiNAGAO/go_gin_gorm_crud_api/initializers"
	"github.com/YuichiNAGAO/go_gin_gorm_crud_api/models"
	"github.com/gin-gonic/gin"
)

func PostsCreate(c *gin.Context) {
	// リクエストボディを受け取る
	var body struct {
		Title string
		Body  string
	}
	c.Bind(&body)

	// ポストを作成する
	var post models.Post

	post.Body = body.Body
	post.Title = body.Title

	result := initializers.DB.Create(&post)

	if result.Error != nil {
		c.JSON(500, gin.H{"error": result.Error})
		return
	}

	c.JSON(200, post)
}
