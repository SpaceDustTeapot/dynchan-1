package models

import "github.com/jinzhu/gorm"

// User represents end users of the system.
type User struct {
	gorm.Model

	UserName    string `gorm:"column:username;size:255;unique"`
	DisplayName string `gorm:"column:name;size:255;null"`
	Email       string `gorm:"column:email;size:255;null"`

	DefaultInsertDisplayName bool `gorm:"column:default_insert_name"`
	DefaultInsertEmail       bool `gorm:"column:default_insert_email"`

	PublicKeys     []UserPublicKey
	ListPublicKeys bool `gorm:"column:list_keys"`
}

// TableName returns a stable table name string for the User entity.
func (User) TableName() string {
	return "users"
}
