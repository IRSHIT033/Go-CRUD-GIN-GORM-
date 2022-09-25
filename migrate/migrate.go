package main

import (
	 "GIN/models"
	"GIN/initializers"
)

func init() {
	initializers.LoadEnvVar()
	initializers.ConnectDB()
}

func main(){
	initializers.DB.AutoMigrate(&models.Post{})
}