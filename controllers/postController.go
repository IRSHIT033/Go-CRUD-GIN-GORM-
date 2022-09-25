package controllers

import (
	"GIN/initializers"
	"GIN/models"
    "github.com/gin-gonic/gin"
)

func PostsCreate (c * gin.Context){
	//get data off req body
	var body struct {
		Title string
		Body string
	}
	c.Bind(&body)


	//Create a post
	post:= models.Post{Title:body.Title,Body:body.Body}
    result := initializers.DB.Create(&post)
    if result.Error != nil {
		c.Status(400)
		return
	}

	c.JSON(200,gin.H{
	 "post":post,
	})
 }

func PostsIndex(c *gin.Context){
	//Get the posts
	var posts []models.Post
	initializers.DB.Find(&posts)
    
	//Respond  with them
    c.JSON(200,gin.H{
		"post":posts,
	})
} 

func PostsShow(c *gin.Context){
	//Get id off url
	id := c.Param("id")
	//Get the posts
	var post []models.Post
	initializers.DB.First(&post,id)

	//respond with them
    c.JSON(200,gin.H{
		"post":post,
	})
}

func PostsUpdate(c *gin.Context){
	//Get the id off the url
	id:=c.Param("id")
	//Get the data off req body
	var body struct {
		Title string
		Body string
	}
	c.Bind(&body)

	// Find the post were uodating
    var post models.Post
	initializers.DB.First(&post,id)

	//update it 
    initializers.DB.Model(&post).Updates(models.Post{
		Title: body.Title,
		Body: body.Body,
	})

	//respond with it
	c.JSON(200,gin.H{
		"posts":post,
	})
}

func PostDelete(c *gin.Context){
    //Get the id off the url
	id:=c.Param("id")
	//Delete the post
	initializers.DB.Delete(&models.Post{},id)
	//respond with it
	c.Status(200)
}