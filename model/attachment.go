package model

type Attachment struct {
	ID       uint64 `gorm:"column:id;primary_key;unique;not null"`
	User     *User  `gorm:"column:user;null"`
	MimeType string `gorm:"column:mime;type:varchar(255);null"`
	FileName string `gorm:"column:filename;type:varchar(255);null"`
}
