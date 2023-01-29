package main

import (
	"github.com/siddhuR/go-CRUD/initializers"
	"github.com/siddhuR/go-CRUD/models"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

func main() {
	initializers.DB.AutoMigrate(&models.POST{})
}
