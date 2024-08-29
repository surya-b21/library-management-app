package model

import "github.com/google/uuid"

type Book struct {
	Base
	BookAPI
}

type BookAPI struct {
	Title      *string    `gorm:"type:varchar(100)"`
	Pages      *int32     `gorm:"type:integer"`
	AuthorID   *uuid.UUID `gorm:"type:varchar(36)"`
	CategoryID *uuid.UUID `gorm:"type:varchar(36)"`
}
