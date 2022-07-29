package main

import (
	"gorm-gin-practise/controllers"
	"gorm-gin-practise/initializers"

	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	_ "gorm-gin-practise/docs"
)

func init() {
	initializers.LoadEnvVariabales()
	initializers.ConnectToDB()
}

// @title           ToDo app
// @version         1.0
// @description     This is a ToDo app in golang(gin, GORM).
// @termsOfService  http://swagger.io/terms/

// @contact.name   API Support
// @contact.url    http://google.com
// @contact.email  asliddintukhtasinov5@gmail.com

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host      localhost:8080
// @BasePath  /api/v1

// @securityDefinitions.apikey  ApiKeyAuth
// @in                          header
// @name                        Authorization
// @description					Description for what is this security definition being used
func main() {

	r := gin.Default()
	v1 := r.Group("/api/v1")
	{
		todo := v1.Group("/")
		{
			todo.POST("/", controllers.ToDoCreate)
			todo.GET("/", controllers.ToDoList)
			todo.POST("/post-active-update/:id", controllers.ToDoUpdateActive)
			todo.GET("/:id", controllers.ToDoDetail)
			todo.PUT("/:id", controllers.ToDoUpdate)
			todo.DELETE("/:id", controllers.ToDoDelete)
		}
	}
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	r.Run(":8080") // listen and serve on 0.0.0.0:8080

}
