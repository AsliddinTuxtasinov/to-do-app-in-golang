package models

import (
	"gorm.io/gorm"
)

type ToDos []ToDo

type ToDo struct {
	gorm.Model
	Title    string `json:"title"`
	Body     string `json:"body"`
	IsActive bool   `json:"is_active" gorm:"type:bool;default:true"`
}
