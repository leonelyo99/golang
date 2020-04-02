package models

import (
	"github.com/jinzhu/gorm"
)

//comment comentario del sistema
type Comment struct{
	gorm.Model
	UserID uint8 `json: "user_id" gorm:"not null; unique"`
	User User `json: "-"`
	Content string `json: "content" gorm:"not null; unique"`
}