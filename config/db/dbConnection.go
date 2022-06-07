package db

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func DbOpen() (*gorm.DB, error) {
	// 詳細は https://github.com/go-sql-driver/mysql#dsn-data-source-name を参照
	dsn := "root:kingdom815@tcp(127.0.0.1:3306)/omen?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Printf("db connection error: %v", err)
		return nil, err
	}
	return db, nil
}
