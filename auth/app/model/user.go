package model

import "github.com/go-openapi/strfmt"

// User struct
type User struct {
	Base
}

// UserAPI struct
type UserAPI struct {
	Name     *string          `json:"name,omitempty" gorm:"type:varchar(50)"`
	Username *string          `json:"username,omitempty" gorm:"type:varchar(50)"`
	Password *strfmt.Password `json:"-" gorm:"type:text"`
}
