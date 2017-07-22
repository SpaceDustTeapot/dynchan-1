package models

// ImageBoard represents a Board for shitposting.
type ImageBoard struct {
	ID             int  `gorm:"column:id;primary_key"`
	PermitTripCode bool `gorm:"column:permit_trip"`
}

// TableName returns a stable table name string for the ImageBoard entity.
func (ImageBoard) TableName() string {
	return "image_boards"
}
