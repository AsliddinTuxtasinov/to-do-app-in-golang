package main

import (
	"fmt"
	"gorm-gin-practise/initializers"
	"gorm-gin-practise/models"
	"log"
)

func init() {
	initializers.LoadEnvVariabales()
	initializers.ConnectToDB()
}

func main() {
	err := initializers.DB.AutoMigrate(&models.ToDo{})
	if err != nil {
		log.Fatal(err)
	}

	err = initializers.DB.AutoMigrate(&models.User{})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Migrated models done ...")
}
