package main

import (
	"github.com/mwolfhoffman/go-crud-api/initializers"
	"github.com/mwolfhoffman/go-crud-api/models"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDb()
}

func main() {
	initializers.DB.AutoMigrate(&models.Post{})
}
