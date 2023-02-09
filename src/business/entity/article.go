package entity

import "gorm.io/gorm"

type Article struct {
	gorm.Model
	Title string
	Body  string
}

type ArticleParam struct {
	Id uint `uri:"id" binding:"required"`
}

type ArticleInputParam struct {
	Tilte string `json:"title" binding:"required"`
	Body  string `json:"body" binding:"required"`
}
