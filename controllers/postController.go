package controllers

import (
	"gorm-gin-practise/initializers"
	"gorm-gin-practise/models"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ToDoCreateUtils struct {
	Title string `json:"title"`
	Body  string `json:"body"`
}

// Create ToDo godoc
// @Summary      Create Todo
// @Description  create by json Todo
// @Tags         ToDo
// @Accept       json
// @Produce      json
// @Param        ToDo body ToDoCreateUtils  true  "Create ToDo"
// @Success      200  {object} models.ToDo
// @Router       / [post]
func ToDoCreate(c *gin.Context) {
	// Get data of request body
	var body ToDoCreateUtils
	err := c.BindJSON(&body)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Status Bad Request"})
		return
	}

	// Create a ToDo
	reqUser, _ := c.Get("user")
	toDo := models.ToDo{Title: body.Title, Body: body.Body, UserID: reqUser.(models.User).ID}
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

// List ToDo godoc
// @Summary      List ToDos
// @Description  get ToDos list
// @Tags         ToDo
// @Accept       json
// @Produce      json
// @Param        q   query    int false "is_active filter by q" Format(int64)
// @Success      200 {object} models.ToDo
// @Router       / [get]
func ToDoList(c *gin.Context) {
	// Get the param
	q := c.Query("q")

	// Get the posts
	var toDos models.ToDos
	reqUser, _ := c.Get("user")
	switch q {
	case "1":
		initializers.DB.Where("is_active = ? AND user_id = ?", true, reqUser.(models.User).ID).Find(&toDos)
		break
	case "0":
		initializers.DB.Where("is_active = ? AND user_id = ?", false, reqUser.(models.User).ID).Find(&toDos)
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

// Detail ToDo godoc
// @Summary      Detail ToDo
// @Description  get ToDo by ID
// @Tags         ToDo
// @Accept       json
// @Produce      json
// @Param        id   path      int  true  "ToDo ID"
// @Success      200  {object}  models.ToDo
// @Router       /{id} [get]
func ToDoDetail(c *gin.Context) {
	// Get "id" off url
	id := c.Param("id")

	// Get the posts
	var toDo models.ToDo
	reqUser, _ := c.Get("user")
	// initializers.DB.First(&toDo, id)
	initializers.DB.Where("user_id = ?", reqUser.(models.User).ID).First(&toDo, id)

	// Response them
	c.JSON(http.StatusOK, gin.H{
		"content": toDo,
	})
}

// Update ToDo godoc
// @Summary      Update ToDo
// @Description  Update ToDo by ID
// @Tags         ToDo
// @Accept       json
// @Produce      json
// @Param        id   path int true "ToDo ID" Format(int64)
// @Param        ToDo body ToDoCreateUtils true "Create ToDo"
// @Success      200  {object} models.ToDo
// @Router       /{id} [put]
func ToDoUpdate(c *gin.Context) {
	// Get "id" off url
	id := c.Param("id")

	// Get the data off request body
	var body ToDoCreateUtils
	err := c.ShouldBindJSON(&body)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Status Bad Request"})
		return
	}

	// Find the ToDo were updating
	var toDo models.ToDo
	reqUser, _ := c.Get("user")
	// initializers.DB.First(&toDo, id)
	initializers.DB.Where("user_id = ?", reqUser.(models.User).ID).First(&toDo, id)

	// Update it
	initializers.DB.Model(&toDo).Updates(models.ToDo{Title: body.Title, Body: body.Body})

	// Respond it
	c.JSON(http.StatusOK, gin.H{
		"content": toDo,
	})
}

// Update ToDo Active godoc
// @Summary      Update ToDo Active
// @Description  Update ToDo Active by ID
// @Tags         ToDo
// @Accept       json
// @Produce      json
// @Param        id  path     int true "ToDo ID" Format(int64)
// @Success      200 {object} models.ToDo
// @Router       /post-active-update/{id} [post]
func ToDoUpdateActive(c *gin.Context) {
	// Get "id" off url
	id := c.Param("id")

	// Find the ToDo were updating
	var toDo models.ToDo
	reqUser, _ := c.Get("user")
	// initializers.DB.First(&toDo, id)
	initializers.DB.Where("user_id = ?", reqUser.(models.User).ID).First(&toDo, id)

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

// Delete ToDo godoc
// @Summary      Delete ToDo
// @Description  Delete ToDo by ID
// @Tags         ToDo
// @Accept       json
// @Produce      json
// @Param        id   path      int  true  "ToDo ID"  Format(int64)
// @Success      204  {string}  http.StatusOK
// @Router       /{id} [delete]
func ToDoDelete(c *gin.Context) {
	// Get "id" off url
	id := c.Param("id")

	// Delete the content
	reqUser, _ := c.Get("user")
	// initializers.DB.Delete(&models.ToDo{}, id)
	tx := initializers.DB.Where("user_id = ?", reqUser.(models.User).ID).Delete(&models.ToDo{}, id)
	if tx.Error != nil {
		log.Fatal(">>>> ", tx.Error)
	}

	// Respond it
	c.JSON(http.StatusOK, gin.H{
		"message": "Deleted Content",
	})
}
