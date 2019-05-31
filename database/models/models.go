package models

import (
	"fmt"

	"github.com/jinzhu/gorm"
)

// User model
type User struct {
	gorm.Model
	Username string `gorm:"size:190;not null"`
	Email    string `gorm:"size:190;unique_index;not null"`
	Password string `json:",omitempty" gorm:"size:190;not null"`
	Age      int
}

// Migrate automigration
func Migrate(db *gorm.DB) {
	db.AutoMigrate(&User{})
	fmt.Println("Tables migrated")
}
