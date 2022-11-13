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

func AddProduct(product *Product) error {
	conn, err := GetDbConnection()
	if err != nil {
		return err
	}
	conn.Create(product)
	return nil
}

func GetProductById(id *uint) (*Product, error) {
	conn, err := GetDbConnection()
	if err != nil {
		return nil, err
	}
	product := &Product{}
	result := conn.First(&product, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return product, nil
}

func GetProducts() ([]Product, error) {
	conn, err := GetDbConnection()
	if err != nil {
		return nil, err
	}
	var products []Product
	result := conn.Find(&products)
	if result.Error != nil {
		return nil, result.Error
	}
	return products, nil
}
