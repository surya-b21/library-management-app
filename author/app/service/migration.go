package service

import (
	"github.com/surya-b21/library-management-app/author/app/model"

	"gorm.io/gorm"
)

// AutoMigrate function
func AutoMigrate(db *gorm.DB) error {
	return db.AutoMigrate(ModelList...)
}

// ModelList list of model
var ModelList []interface{} = []interface{}{
	&model.Author{},
}
