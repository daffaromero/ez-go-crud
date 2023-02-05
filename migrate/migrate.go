package main

import (
	"github.com/daffaromero/go-crud/initializers"
	"github.com/daffaromero/go-crud/models"
)

func init() {
	initializers.LoadEnvVars()
	initializers.ConnectToDB()
}

func main() {
	initializers.DB.AutoMigrate(&models.Post{})
}
