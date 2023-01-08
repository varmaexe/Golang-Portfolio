package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	// ID       string `gorm:"SERIAL PRIMAY KEY"`
	Email    string `gorm:"unique"`
	Password string
}
