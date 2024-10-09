package models

import "gorm.io/gorm"

// User struct defines the user model
type User struct {
	gorm.Model
	Name  string `json:"name" gorm:"size:100;not null"`
	Email string `json:"email" gorm:"unique;not null"`
	Age   int    `json:"age"`
}
