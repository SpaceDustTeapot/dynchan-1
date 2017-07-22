package model

// PostAttachment represents assocations between Attachment and Post entities.
type PostAttachment struct {
	ID   uint64    `gorm:"column:id;primary_key;unique;not null"`
	Post Post      `gorm:"column:post"`
	Key  PublicKey `gorm:"column:key"`
}

// TableName returns a stable table name string for the PostAttachment entity.
func (PostAttachment) TableName() string {
	return "post_attachments"
}
