package model

import (
	"time"

	"golang.org/x/crypto/openpgp/packet"

	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

// PublicKey is a model for PGP public keys associated with users, post and DM signing, post and DM encryption,
// and board-level web-of-trust networks.
type PublicKey struct {
	Fingerprint [20]byte                  `gorm:"column:fingerprint;primary_key;unique;not null;type:binary(20)"`
	KeyID       *uint64                   `gorm:"column:key_id;index:idx_key_id;null"`
	Algorithm   packet.PublicKeyAlgorithm `gorm:"column:algorithm;not null;type:tinyint"`
	CreatedAt   time.Time                 `gorm:"column:created;not null"`
	UpdatedAt   time.Time                 `gorm:"column:updated;not null"`
	DeletedAt   *time.Time                `gorm:"column:deleted;null"`
}

// TableName returns a stable table name string for the PublicKey entity.
func (PublicKey) TableName() string {
	return "public_keys"
}
