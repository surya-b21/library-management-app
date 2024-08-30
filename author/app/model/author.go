package model

type Author struct {
	Base
	AuthorAPI
}

type AuthorAPI struct {
	Name *string `gorm:"type:varchar(100)"`
}
