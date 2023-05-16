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

func PostsIndex(c *gin.Context) {
	var posts []models.Post
	result := initializers.DB.Find(&posts)

	if result.Error != nil {
		c.JSON(500, gin.H{"error": result.Error})
		return
	}

	c.JSON(200, gin.H{"posts": posts})
}

func PostsShow(c *gin.Context) {
	var post models.Post
	result := initializers.DB.First(&post, c.Param("id"))

	if result.Error != nil {
		c.JSON(500, gin.H{"error": result.Error})
		return
	}

	c.JSON(200, post)
}

func PostsUpdate(c *gin.Context) {
	var post models.Post
	result := initializers.DB.First(&post, c.Param("id"))

	if result.Error != nil {
		c.JSON(500, gin.H{"error": result.Error})
		return
	}

	var body struct {
		Title string
		Body  string
	}
	c.Bind(&body)

	// post.Body = body.Body
	// post.Title = body.Title

	// result = initializers.DB.Save(&post)

	initializers.DB.Model(&post).Updates(models.Post{Title: body.Title, Body: body.Body})

	if result.Error != nil {
		c.JSON(500, gin.H{"error": result.Error})
		return
	}

	c.JSON(200, post)
}

func PostsDelete(c *gin.Context) {
	result := initializers.DB.Delete(&models.Post{}, c.Param("id"))
	if result.Error != nil {
		c.JSON(500, gin.H{"error": result.Error})
		return
	}

	c.Status(204)
}
