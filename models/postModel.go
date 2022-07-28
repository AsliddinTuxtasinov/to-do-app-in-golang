package models

import (
	"gorm.io/gorm"
)

type ToDos []ToDo

type ToDo struct {
	gorm.Model
	Title    string
	Body     string
	IsActive bool `json:"is_active" gorm:"type:bool;default:true"`
}
