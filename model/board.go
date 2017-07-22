package model

// Board represents a community or discussion area.
type Board struct {
	ID              uint64 `gorm:"column:id;primary_key;unique;not null"`
	Community       *Board `gorm:"column:parent;null"`
	URI             string `gorm:"column:uri;not null"`
	Type            uint8
	Owner           uint8
	Password        [60]byte `gorm:"column:password;null"`
	PermitClearnet  bool     `gorm:"column:permit_clear"`
	PermitTor       bool     `gorm:"column:permit_tor"`
	PermitUser      bool     `gorm:"column:permit_user"`
	PermitAnonymous bool     `gorm:"column:permit_anon"`
}
