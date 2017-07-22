package models

// Attachment represents a file.
type Attachment struct {
	ID       int    `gorm:"column:id;primary_key"`
	User     *User  `gorm:"column:user;null"`
	MimeType string `gorm:"column:mime;type:varchar(255);null"`
	FileName string `gorm:"column:filename;type:varchar(255);null"`
}

// TableName returns a stable table name string for the Attachment entity.
func (Attachment) TableName() string {
	return "attachments"
}
