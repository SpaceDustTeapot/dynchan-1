package model

// ImageBoard represents a Board for shitposting.
type ImageBoard struct {
	ID             uint64
	PermitTripCode bool `gorm:"column:permit_trip"`
}

// TableName returns a stable table name string for the ImageBoard entity.
func (ImageBoard) TableName() string {
	return "image_boards"
}
