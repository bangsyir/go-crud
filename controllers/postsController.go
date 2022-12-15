package controllers

import (
	"net/http"

	"github.com/bangsyir/go-crud/initializers"
	"github.com/bangsyir/go-crud/models"
	"github.com/gin-gonic/gin"
)

func PostCreate(c *gin.Context) {
	// get data from request body

	var body struct {
		Title string
		Desc  string
	}

	c.Bind(&body)

	// create a post

	post := models.Post{Title: body.Title, Desc: body.Desc}
	result := initializers.DB.Create(&post)

	if result.Error != nil {
		c.Status(http.StatusBadRequest)
		return
	}
	//return it

	c.JSON(http.StatusCreated, gin.H{
		"post": post,
	})
}

func PostsIndex(c *gin.Context) {
	// get post
	var posts []models.Post
	initializers.DB.Find(&posts)

	//response
	c.JSON(http.StatusOK, gin.H{
		"posts": posts,
	})
}

func PostShow(c *gin.Context) {
	//get id from params
	id := c.Param("id")
	var post models.Post
	initializers.DB.First(&post, id)

	c.JSON(http.StatusOK, gin.H{
		"post": post,
	})
}

func PostUpdate(c *gin.Context) {
	// get id
	id := c.Param("id")
	// get data req body
	var body struct {
		Title string
		Desc  string
	}
	c.Bind(&body)
	//find post
	var post models.Post
	initializers.DB.First(&post, id)
	//update it
	initializers.DB.Model(&post).Updates(models.Post{
		Title: body.Title,
		Desc:  body.Desc,
	})
	// response
	c.JSON(http.StatusOK, gin.H{
		"post": post,
	})
}

func PostDelete(c *gin.Context) {
	// get params
	id := c.Param("id")
	//delete post
	initializers.DB.Delete(&models.Post{}, id)
	//response
	c.JSON(http.StatusAccepted, gin.H{
		"message": "post is deleted",
	})
}
