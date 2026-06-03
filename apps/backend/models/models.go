package models

import "gorm.io/gorm"

type Task struct {
	gorm.Model
	TaskName string		`json:"task"`
	Description  string		`json:"description"`
}