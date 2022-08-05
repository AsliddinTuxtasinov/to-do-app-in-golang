package main

import (
	"fmt"
	"gorm-gin-practise/initializers"
	"gorm-gin-practise/models"
)

func init() {
	initializers.LoadEnvVariabales()
	initializers.ConnectToDB()
}

func main() {
	initializers.DB.AutoMigrate(&models.ToDo{}, &models.User{})
	fmt.Println("Migrated models done ...")
}
