package database

import (
	"time"
)

type Product struct {
	ID uint `json:"id" gorm:"primaryKey"`
	// UUID        uuid.UUID `json:"uuid" gorm:"type:uuid;default:uuid_generate_v4()"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"createdAt" gorm:"autoCreateTime"`
	UpdatedAt   time.Time `json:"updatedAt" gorm:"autoUpdateTime"`
	ClientID    uint      `json:"clientId"`
	Client      Client
}

func NewProduct() *Product {
	return &Product{}
}

func CreateProduct(newProduct *Product) error {
	conn, err := GetDbConnection()
	if err != nil {
		return err
	}
	conn.Create(newProduct)
	return nil
}
