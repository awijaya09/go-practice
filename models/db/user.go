package db

import "github.com/jinzhu/gorm"

// User struct for DB Object
type User struct {
	gorm.Model
	FirstName string
	LastName  string
	Email     string `gorm:"type:varchar(100);unique_index"`
	Password  string
}
