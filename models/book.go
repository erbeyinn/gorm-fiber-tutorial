package models

import "gorm.io/gorm"

type Book struct {
	gorm.Model
	Id *int `json:"id"`
	Author string `json:"author"`
	BookName string `json:"book_name"`
}