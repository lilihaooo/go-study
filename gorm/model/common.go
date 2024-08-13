package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username    string
	Age         int
	Nationality string
	Tags        []Tag `gorm:"many2many:users_tags;"`
	Books       []Book
}

type Tag struct {
	ID      uint
	TagName string
}

type Book struct {
	gorm.Model
	BookName string
	UserID   uint
}
