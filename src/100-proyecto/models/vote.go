package models

import (
	"github.com/jinzhu/gorm"
)

//permite controlar que un usuairo solo vote una unica vez por cada comentario
type Vote struct{
	gorm.Model
	CommentId uint `json: "commentId" gorm:"not null"`
	UserID uint `json:"userId" gorm:"not null"`
	Value bool `json:"value" gorm:"not null"`
}