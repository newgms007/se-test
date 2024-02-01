package entity

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username string `valid:"required~Username is required"`
	Email    string `valid:"required~Email is required, email~Email is invalid"`
}
