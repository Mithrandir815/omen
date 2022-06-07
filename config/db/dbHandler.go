package db

import (
	"fmt"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name   string
	Weight int
	Height int
}

func Migrate(db *gorm.DB) {
	var user User
	err := db.AutoMigrate(&user)
	if err != nil {
		fmt.Printf("Migrate error: %v", err)
	}
}
