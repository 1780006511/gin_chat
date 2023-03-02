package models

import "gorm.io/gorm"

type Contact struct {
	gorm.Model
}

func (table *Contact) TableName() string {
	return "contact"
}
