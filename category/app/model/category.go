package model

type Category struct {
	Base
	CategoryAPI
}

type CategoryAPI struct {
	Name *string `gorm:"type:varchar(100)"`
}
