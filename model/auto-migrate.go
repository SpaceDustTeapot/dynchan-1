package model

import "github.com/jinzhu/gorm"

// AutoMigrate executes the database auto-migration system from Gorm over each entity defined in the model package.
func AutoMigrate(db *gorm.DB) *gorm.DB {
	return db.AutoMigrate(
		&Attachment{},
		&BoardBanner{},
		&Board{},
		&ImageBoard{},
		&PostAttachment{},
		&PostSignatureKey{},
		&PostTargetKey{},
		&Post{},
		&PublicKey{},
		&StoreBoard{},
		&UserPublicKey{},
		&User{})
}
