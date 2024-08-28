package model

import "github.com/google/uuid"

type Book struct {
	Base
	Title      *string    `gorm:"type:varchar(100)"`
	Pages      *int       `gorm:"type:integer"`
	AuthorID   *uuid.UUID `gorm:"type:varchar(36)"`
	CategoryID *uuid.UUID `gorm:"type:varchar(36)"`
}
