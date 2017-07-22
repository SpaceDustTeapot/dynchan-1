package model

// PostSignatureKey represents signing keys obtained from the metadata of a PGP signed Post.
// It is a direct reimplementation of the PostTargetKey entity.
type PostSignatureKey struct {
	Post Post      `gorm:"column:post;primary_key"`
	Key  PublicKey `gorm:"column:key;primary_key"`
}

// TableName is returns a stable table name string for the PostSignatureKey entity.
func (PostSignatureKey) TableName() string {
	return "post_signature_keys"
}
