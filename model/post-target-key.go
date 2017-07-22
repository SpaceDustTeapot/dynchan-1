package model

// PostTargetKey represents Post encryption keys obtained from the metadata of a PGP encrypted Post.
type PostTargetKey struct {
	Post Post      `gorm:"column:post;primary_key"`
	Key  PublicKey `gorm:"column:key;primary_key"`
}

// TableName returns a stable table name string for the PostTargetKey entity.
func (PostTargetKey) TableName() string {
	return "post_target_keys"
}
