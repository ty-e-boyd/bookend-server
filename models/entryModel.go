package models

import "gorm.io/gorm"

type Tags []string

type Entry struct {
	gorm.Model
	Title      string
	Body       string
	BookName   string
	Author     string
	PageNumber string
}
