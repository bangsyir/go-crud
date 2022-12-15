package main

import (
	"github.com/bangsyir/go-crud/controllers"
	"github.com/bangsyir/go-crud/initializers"
	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvVariable()
	initializers.ConnectToDB()

}

func main() {
	r := gin.Default()

	r.POST("/posts", controllers.PostCreate)
	r.GET("/posts", controllers.PostsIndex)
	r.GET("/posts/:id", controllers.PostShow)
	r.PUT("/posts/:id", controllers.PostUpdate)
	r.DELETE("/posts/:id", controllers.PostDelete)

	r.Run()
}
