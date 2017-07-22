package model

type UserPublicKey struct {
	User   User      `gorm:"column:user;not null;index"`
	Key    PublicKey `gorm:"column:key;not null"`
	Active bool      `gorm:"column:active"`
}

// TableName returns a stable table name string for the UserPublicKey entity.
func (UserPublicKey) TableName() string {
	return "user_public_keys"
}
