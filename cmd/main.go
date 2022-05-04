package main

import (
	"fmt"
	"project/omen/config/db"

	"gorm.io/gorm"
)

func main() {
	var data *gorm.DB
	data, err := db.DbOpen()
	if err != nil {
		fmt.Printf("DB connection error: %v", err)
	}
	db.Migrate(data)
}
