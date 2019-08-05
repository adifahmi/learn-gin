package models

import (
	"fmt"

	"github.com/jinzhu/gorm"
)

// Migrate automigration
func Migrate(db *gorm.DB) {
	db.AutoMigrate(&User{})
	fmt.Println("Tables migrated")
}
