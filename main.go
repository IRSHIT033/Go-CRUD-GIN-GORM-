package main

import (
	"GIN/initializers"
    "github.com/gin-gonic/gin"
	"GIN/controllers"
)

func init(){
	initializers.LoadEnvVar()
	initializers.ConnectDB()
}

func main() {
	r:=gin.Default()
	//r.GET("/",controllers.PostsCreate)
	r.POST("/post",controllers.PostsCreate)
	r.GET("/getPosts",controllers.PostsIndex)
	r.GET("/posts/:id",controllers.PostsShow)
	r.PUT("/UpdatePost/:id",controllers.PostsUpdate)
	r.DELETE("/deletePost/:id",controllers.PostDelete)
	r.Run()
}