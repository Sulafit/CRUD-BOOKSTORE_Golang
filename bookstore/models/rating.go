package models

import "gorm.io/gorm"

type Rating struct {
	gorm.Model
	UserID int
	BookID int
	Value  int
}
