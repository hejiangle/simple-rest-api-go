package dto

import "github.com/jinzhu/gorm"

type TodoItem struct {
	gorm.Model
	Status bool
	Content string
}