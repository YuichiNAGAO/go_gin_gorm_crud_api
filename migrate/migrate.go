package main

import (
	"github.com/YuichiNAGAO/go_gin_gorm_crud_api/initializers"
	"github.com/YuichiNAGAO/go_gin_gorm_crud_api/models"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectDB()
}

func main() {
	initializers.DB.AutoMigrate(&models.Post{})
}
