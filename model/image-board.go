package model

type ImageBoard struct {
	ID             uint64
	PermitTripCode bool `gorm:"column:permit_trip"`
}

func (ImageBoard) TableName() string {
	return "image_boards"
}
