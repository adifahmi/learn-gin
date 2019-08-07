package models

import (
	"github.com/adifahmi/learn-gin/lib"
	"github.com/jinzhu/gorm"
)

// User model
type User struct {
	gorm.Model `json:"-"`
	Username   string `gorm:"size:190;not null"`
	Email      string `gorm:"size:190;unique_index;not null"`
	Password   string `json:",omitempty" gorm:"size:190;not null"`
	Age        int
}

// Serialize serializes user data
func (u *User) Serialize() lib.JSON {
	return lib.JSON{
		"id":       u.ID,
		"username": u.Username,
		"email":    u.Email,
		"age":      u.Age,
	}
}
