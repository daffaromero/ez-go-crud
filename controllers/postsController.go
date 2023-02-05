package controllers

import (
	"github.com/daffaromero/go-crud/initializers"
	"github.com/daffaromero/go-crud/models"
	"github.com/gin-gonic/gin"
)

func PostsCreate(c *gin.Context) {
	// Get data from req body
	var body struct {
		Body  string
		Title string
	}

	c.Bind(&body)

	// Create a post
	post := models.Post{Title: body.Title, Body: body.Body}
	result := initializers.DB.Create(&post)

	if result.Error != nil {
		c.Status(400)
		return
	}

	// Return the post
	c.JSON(200, gin.H{
		"post": post,
	})
}

func PostsIndex(c *gin.Context) {
	// Get the posts
	var posts []models.Post
	initializers.DB.Find(&posts)

	// Respond with the posts
	c.JSON(200, gin.H{
		"post": posts,
	})
}

func PostsShow(c *gin.Context) {
	// Get ID from URL
	id := c.Param("id")

	// Get the posts
	var post models.Post
	initializers.DB.First(&post, id)

	// Respond with the posts
	c.JSON(200, gin.H{
		"post": post,
	})
}

func PostsUpdate(c *gin.Context) {
	// Get ID from URL
	id := c.Param("id")

	// Get data from req body
	var body struct {
		Body  string
		Title string
	}

	c.Bind(&body)

	// Find post to update
	var post models.Post
	initializers.DB.First(&post, id)

	// Update it
	initializers.DB.Model(&post).Updates(models.Post{
		Title: body.Title,
		Body:  body.Body,
	})

	// Respond with the update
	c.JSON(200, gin.H{
		"post": post,
	})
}

func PostsDelete(c *gin.Context) {
	// Get ID from URL
	id := c.Param("id")

	// Delete the post
	initializers.DB.Delete(&models.Post{}, id)

	// Respond
	c.Status(200)
}
