package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/vermaexe/go-jwt/initializers"
	"github.com/vermaexe/go-jwt/models"
)

func PostsCreate(c *gin.Context) {
	//get data from req body
	var body struct {
		Body  string `form:"Body"`
		Title string `form:"Title"`
	}
	c.Bind(&body)
	//create post
	post := models.Post{Title: body.Title, Body: body.Body}
	result := initializers.DB.Create(&post)
	if result.Error != nil {
		c.Status(400)
		return
	}

	//return it
	// c.JSON(200, gin.H{"post": post})
	c.HTML(200, "createposts.html", post)
}
func PostsIndex(c *gin.Context) {
	//gets all the posts
	var posts []models.Post
	initializers.DB.Find(&posts)

	//respond with them
	// c.JSON(200, gin.H{"post": posts})
	c.HTML(200, "posts.html", posts)
}
func PostShow(c *gin.Context) {
	//get id from url
	// id := c.Param("id")
	//get id from form value
	id := c.Request.FormValue("id")
	//gets all the posts
	var post models.Post
	initializers.DB.First(&post, id)

	//respond with them
	// c.JSON(200, gin.H{"post": post})

	//respond with them to frontend
	c.HTML(200, "posts.html", post)
}
func PostUpdate(c *gin.Context) {
	//get the id off the url
	// id := c.Param("id")
	id := c.Request.FormValue("id")

	//get the data off req body
	var body struct {
		Body  string `form:"Body"`
		Title string `form:"Title"`
	}
	c.Bind(&body)

	//find the post we want to updte
	var post models.Post
	initializers.DB.Find(&post, id)

	//update it
	initializers.DB.Model(&post).Updates(models.Post{
		Title: body.Title,
		Body:  body.Body,
	})
	//respond with it
	c.HTML(200, "updatepost.html", post)
}
func PostsDelete(c *gin.Context) {
	//get the id from url
	id := c.Request.FormValue("id")

	//delete the posts
	initializers.DB.Delete(&models.Post{}, id)

	//respond
	c.HTML(200, "deletepost.html", gin.H{})
}
