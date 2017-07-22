package model

// BoardBanner represents attachment, title, and subtitle banners for Board entities.
type BoardBanner struct {
	ID         uint64      `gorm:"column:id;primary_key;unique;not null"`
	Board      Board       `gorm:"column:board;not null"`
	Attachment *Attachment `gorm:"column:attachment;null"`
	Title      string      `gorm:"column:title;type:varchar(255);null"`
	SubTitle   string      `gorm:"column:subtitle;type:varchar(2048);null"`
}

// TableName returns a stable table name string for the BoardBanner entity.
func (BoardBanner) TableName() string {
	return "board_banners"
}
