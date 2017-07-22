package model

type BoardBanner struct {
	ID         uint64      `gorm:"column:id;primary_key;unique;not null"`
	Board      Board       `gorm:"column:board;not null"`
	Attachment *Attachment `gorm:"column:attachment;null"`
	Title      string      `gorm:"column:title;type:varchar(255);null"`
	SubTitle   string      `gorm:"column:subtitle;type:varchar(2048);null"`
}

func (BoardBanner) TableName() string {
	return "board_banners"
}
