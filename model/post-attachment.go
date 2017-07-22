package model

// PostAttachment is a model for associating Attachment entities with a Post.
type PostAttachment struct {
	ID   uint64    `gorm:"column:id;primary_key;unique;not null"`
	Post Post      `gorm:"column:post"`
	Key  PublicKey `gorm:"column:key"`
}

// TableName is a function that returns a stable table name string for the PostAttachment entity.
func (PostAttachment) TableName() string {
	return "post_attachments"
}
