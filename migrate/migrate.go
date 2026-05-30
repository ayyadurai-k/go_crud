package main

import (
	"crud/initializers"
	"crud/models"
)

func init()  {
	initializers.LoadEnvVariables()
	initializers.ConnectDB()
}

func main(){
	initializers.DB.AutoMigrate(&models.Post{})
}