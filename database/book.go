package database

import (
	"time"
)

type Book struct {
	ID          uint      `json:"id" gorm:"primaryKey"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"createdAt" gorm:"autoCreateTime"`
	UpdatedAt   time.Time `json:"updatedAt" gorm:"autoUpdateTime"`
}

func NewBook() *Book {
	return &Book{}
}

func CreateBook(newBook *Book) error {
	conn, err := GetDbConnection()
	if err != nil {
		return err
	}
	conn.Create(newBook)
	return nil
}
