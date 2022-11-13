package database

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var dbConn *gorm.DB

func CreateDbConnection() error {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	dbConn = db

	db.AutoMigrate(&User{})
	db.AutoMigrate(&Client{})
	db.AutoMigrate(&Book{})
	db.AutoMigrate(&Product{})
	db.AutoMigrate(&Item{})
	db.AutoMigrate(&StockSite{})

	return err
}

func GetDbConnection() (*gorm.DB, error) {
	return dbConn, nil
}
