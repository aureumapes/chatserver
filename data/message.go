package data

import "gorm.io/gorm"

type Message struct {
	Chat    string
	Author  string
	Content string
	Id      string
	gorm.Model
}
