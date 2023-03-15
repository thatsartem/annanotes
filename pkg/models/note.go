package models

import "gorm.io/gorm"

type Note struct {
	gorm.Model
	Title  string `json:"title"`
	Text   string `json:"text"`
	UserID string `json:"user_id"`
}
