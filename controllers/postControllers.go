package controllers

import (
	"github.com/YuichiNAGAO/go_gin_gorm_crud_api/initializers"
	"github.com/YuichiNAGAO/go_gin_gorm_crud_api/models"
	"github.com/gin-gonic/gin"
)

func PostsCreate(c *gin.Context) {
	// リクエストボディを受け取る

	// ポストを作成する

	var post models.Post
	post.Body = "test"
	post.Title = "test"

	// c.BindJSON(&post)
	result := initializers.DB.Create(&post)

	if result.Error != nil {
		c.JSON(500, gin.H{"error": result.Error})
		return
	}

	c.JSON(200, post)
}
