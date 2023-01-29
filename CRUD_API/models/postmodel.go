package models

import "gorm.io/gorm"

type POST struct {
	gorm.Model
	Title string
	Body  string
}
