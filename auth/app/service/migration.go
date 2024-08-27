package service

import (
	"auth/app/model"

	"gorm.io/gorm"
)

// AutoMigrate function
func AutoMigrate(db *gorm.DB) error {
	return db.AutoMigrate(ModelList...)
}

// ModelList list of model
var ModelList []interface{} = []interface{}{
	&model.User{},
}
