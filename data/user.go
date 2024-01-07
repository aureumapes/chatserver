package data

import (
	"gorm.io/gorm"
)

type Role string

const (
	DEV    Role = "Dev"
	ADMIN       = "Administrator"
	MOD         = "Mod"
	MEMBER      = "Member"
)

type User struct {
	Username string
	Token    string
	Role     Role
	Id       string
	gorm.Model
}
