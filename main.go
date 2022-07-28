package main

import (
	"gorm-gin-practise/controllers"
	"gorm-gin-practise/initializers"

	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvVariabales()
	initializers.ConnectToDB()
}

func main() {
	r := gin.Default()
	r.POST("/", controllers.ToDoCreate)
	r.GET("/", controllers.ToDoList)
	r.POST("/post-active-update/:id", controllers.ToDoUpdateActive)
	r.GET("/:id", controllers.ToDoDetail)
	r.PUT("/:id", controllers.ToDoUpdate)
	r.DELETE("/:id", controllers.ToDoDelete)
	r.Run(":8080") // listen and serve on 0.0.0.0:8080
}
