package controllers

import (
	"crud/initializers"
	"crud/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func PostCreate(c *gin.Context) {
	var body struct {
		Title string
		Body  string
	}
	c.Bind(&body)

	post := models.Post{Title: body.Title, Body: body.Body}

	result := initializers.DB.Create(&post) // pass pointer of data to Create

	if result.Error != nil {
		c.Status(400)
		return
	}

	c.JSON(200, gin.H{
		"result": "Post created",
		"postId": post.ID,
	})
}

func PostList(c *gin.Context) {
	var posts []models.Post

	initializers.DB.Find(&posts)

	c.JSON(http.StatusOK, gin.H{
		"posts": posts,
	})
}

func PostDetail(c *gin.Context) {
	var post models.Post

	id := c.Param("id")

	result := initializers.DB.First(&post, id)

	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Post not found",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "OK",
		"post":   post,
	})

}

func PostUpdate(c *gin.Context) {
	id := c.Param("id")

	var body struct {
		Title string
		Body  string
	}
	c.Bind(&body)

	var post models.Post
	initializers.DB.First(&post, id)

	initializers.DB.Model(&post).Updates(models.Post{
		Title: body.Title, Body: body.Body,
	})

	c.JSON(http.StatusOK, gin.H{
		"status": "OK",
		"post":   post,
	})

}

func PostDelete(c *gin.Context) {

	id := c.Param("id")

	initializers.DB.Delete(&models.Post{}, id)

	c.JSON(http.StatusOK, gin.H{
		"status": "OK",
	})

}
