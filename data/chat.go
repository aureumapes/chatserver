package data

import "gorm.io/gorm"

type Chat struct {
	Name string
	Id   string
	gorm.Model
}
