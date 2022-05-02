package db

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func DbConnection() {
	// 詳細は https://github.com/go-sql-driver/mysql#dsn-data-source-name を参照
	dsn := "user:pass@tcp(127.0.0.1:3306)/dbname?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	fmt.Printf("db: %v, err: %v", db, err)
}
