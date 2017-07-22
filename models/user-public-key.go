package models

type UserPublicKey struct {
	ID     int       `gorm:"column:id;primary_key"`
	User   User      `gorm:"column:user"`
	Key    PublicKey `gorm:"column:key"`
	Active bool      `gorm:"column:active"`
}

// TableName returns a stable table name string for the UserPublicKey entity.
func (UserPublicKey) TableName() string {
	return "user_public_keys"
}
