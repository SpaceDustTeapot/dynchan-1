package models

// PostAttachment represents assocations between Attachment and Post entities.
type PostAttachment struct {
	ID   int       `gorm:"column:id;primary_key"`
	Post Post      `gorm:"column:post"`
	Key  PublicKey `gorm:"column:key"`
}

// TableName returns a stable table name string for the PostAttachment entity.
func (PostAttachment) TableName() string {
	return "post_attachments"
}
