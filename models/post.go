package models

import "time"

// Post represents a singular topic contribution to a Board.
type Post struct {
	ID    int   `gorm:"column:id;primary_key"`
	User  *User `gorm:"column:user;null"`
	Board Board

	IsSigned    bool `gorm:"column:is_signed"`
	IsEncrypted bool `gorm:"column:is_encrypted"`

	Caption     string `gorm:"column:caption;size:255;null"`
	DisplayName string `gorm:"column:name;size:255;null"`
	Email       string `gorm:"column:email;size:255;null"`

	Content string `gorm:"column:content"`

	SignatureKeys []PostSignatureKey
	TargetKeys    []PostTargetKey

	CreatedAt time.Time  `gorm:"column:created"`
	UpdatedAt time.Time  `gorm:"column:updated"`
	DeletedAt *time.Time `gorm:"column:deleted"`
}

// TableName returns a stable table name string for the Post entity.
func (Post) TableName() string {
	return "posts"
}
