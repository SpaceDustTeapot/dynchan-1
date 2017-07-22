package models

import "github.com/jinzhu/gorm"

// Board represents a community or discussion area.
type Board struct {
	gorm.Model
	Community *Board `gorm:"column:parent;null" json:"community" description:"parent community"`

	URI      string   `gorm:"column:uri;not null" json:"uri" required:"true" description:"URI mount location"`
	Type     uint8    `gorm:"column:type;not null"`
	Founder  *User    `gorm:"column:founder;null" json:"founder" description:"founding User"`
	Password [60]byte `gorm:"column:password;null" json:"-" description:"password"`

	PermitClearnet  bool `gorm:"column:permit_clear" json:"permitClearNet" description:"permit access from the common Internet"`
	PermitTor       bool `gorm:"column:permit_tor" json:"permitTor" description:"permit access from the Tor network"`
	PermitUser      bool `gorm:"column:permit_user" json:"permitUser" description:"permit access with User accounts"`
	PermitAnonymous bool `gorm:"column:permit_anon" json:"permitAnonymous" description:"permit anonymous access"`
}
