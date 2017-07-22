package model

// StoreBoard is a model for boards consisting of filesystem-like storage shares.
type StoreBoard struct {
	ID uint64

	PermitUserUpload          bool `gorm:"column:permit_user_upload"`
	PermitModeratorUpload     bool `gorm:"column:permit_mod_upload"`
	PermitAdministratorUpload bool `gorm:"column:permit_admin_upload"`
}

// TableName returns a stable table name string for the StoreBoard entity.
func (StoreBoard) TableName() string {
	return "store_boards"
}
