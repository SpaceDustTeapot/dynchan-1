package models

import (
	"time"

	"golang.org/x/crypto/openpgp/packet"

	"database/sql"

	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

// PublicKey represents PGP public keys associated with User entities, Post and DirectMessage signing,
// Post and DirectMessage encryption, and Board entity web-of-trust networks.
type PublicKey struct {
	ID int `gorm:"column:id;primary_key"`

	Fingerprint [20]byte                  `gorm:"column:fingerprint;primary_key;type:binary(20)"`
	KeyID       sql.NullInt64             `gorm:"column:key_id;index:idx_key_id"`
	Algorithm   packet.PublicKeyAlgorithm `gorm:"column:algorithm;not null;type:tinyint"`

	CreatedAt time.Time  `gorm:"column:created;not null"`
	UpdatedAt time.Time  `gorm:"column:updated;not null"`
	DeletedAt *time.Time `gorm:"column:deleted;null"`
}

// TableName returns a stable table name string for the PublicKey entity.
func (PublicKey) TableName() string {
	return "public_keys"
}
