package model

import "time"

// User represents end users of the system.
type User struct {
	Name        string `gorm:"column:name;type:varchar(255);primary_key;not null"`
	DisplayName string `gorm:"column:name;type:varchar(255);null"`
	Email       string `gorm:"column:email;type:varchar(255);null"`

	DefaultInsertDisplayName bool `gorm:"column:default_insert_name"`
	DefaultInsertEmail       bool `gorm:"column:default_insert_email"`

	PublicKeys     []UserPublicKey
	ListPublicKeys bool `gorm:"column:list_keys"`

	CreatedAt time.Time  `gorm:"column:created"`
	UpdatedAt time.Time  `gorm:"column:updated"`
	DeletedAt *time.Time `gorm:"column:created"`
}

// TableName returns a stable table name string for the User entity.
func (User) TableName() string {
	return "users"
}
