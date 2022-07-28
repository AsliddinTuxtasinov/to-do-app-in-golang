package controllers

import (
	"gorm-gin-practise/initializers"
	"gorm-gin-practise/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ToDoCreate(c *gin.Context) {
	// Get data of request body
	var body struct {
		Title string `json:"title"`
		Body  string `json:"body"`
	}
	err := c.BindJSON(&body)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Status Bad Request"})
		return
	}

	// Create a ToDo
	toDo := models.ToDo{Title: body.Title, Body: body.Body}
	result := initializers.DB.Create(&toDo)
	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": result.Error.Error()})
		return
	}

	// Return it
	c.JSON(http.StatusCreated, gin.H{
		"content": toDo,
	})
}

func ToDoList(c *gin.Context) {
	// Get the param
	q := c.Query("q")

	// Get the posts
	var toDos models.ToDos
	switch q {
	case "1":
		initializers.DB.Where("is_active = ?", true).Find(&toDos)
		break
	case "0":
		initializers.DB.Where("is_active = ?", false).Find(&toDos)
		break
	case "":
		initializers.DB.Find(&toDos)
		break
	}

	// Response them
	c.JSON(http.StatusOK, gin.H{
		"contents": toDos,
	})
}

func ToDoDetail(c *gin.Context) {
	// Get "id" off url
	id := c.Param("id")

	// Get the posts
	var toDo models.ToDo
	initializers.DB.First(&toDo, id)

	// Response them
	c.JSON(http.StatusOK, gin.H{
		"content": toDo,
	})
}

func ToDoUpdate(c *gin.Context) {
	// Get "id" off url
	id := c.Param("id")

	// Get the data off request body
	var body struct {
		Title string `json:"title"`
		Body  string `json:"body"`
	}
	err := c.BindJSON(&body)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Status Bad Request"})
		return
	}

	// Find the ToDo were updating
	var toDo models.ToDo
	initializers.DB.First(&toDo, id)

	// Update it
	initializers.DB.Model(&toDo).Updates(models.ToDo{Title: body.Title, Body: body.Body})

	// Respond it
	c.JSON(http.StatusOK, gin.H{
		"content": toDo,
	})
}

func ToDoUpdateActive(c *gin.Context) {
	// Get "id" off url
	id := c.Param("id")

	// Find the ToDo were updating
	var toDo models.ToDo
	initializers.DB.First(&toDo, id)

	// Update it
	if toDo.IsActive == true {
		toDo.IsActive = false
	} else {
		toDo.IsActive = true
	}
	initializers.DB.Save(toDo)

	// Respond it
	c.JSON(http.StatusOK, gin.H{
		"content": toDo,
	})
}

func ToDoDelete(c *gin.Context) {
	// Get "id" off url
	id := c.Param("id")

	// Delete the content
	initializers.DB.Delete(&models.ToDo{}, id)

	// Respond it
	c.JSON(http.StatusOK, gin.H{
		"message": "Deleted Content",
	})
}
