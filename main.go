package main

import (
	"github.com/gin-gonic/gin"
	"github.com/mwolfhoffman/go-crud-api/controllers"
	"github.com/mwolfhoffman/go-crud-api/initializers"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDb()
}

func main() {
	r := gin.Default()

	r.GET("/posts", controllers.GetPosts)
	r.POST("/posts", controllers.CreatePost)
	r.PUT("/posts/:id", controllers.UpdatePost)
	r.GET("/posts/:id", controllers.GetPost)
	r.DELETE("/posts/:id", controllers.DeletePost)

	r.Run()
}
