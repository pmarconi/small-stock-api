package database

import (
	"time"
)

type Item struct {
	ID uint `json:"id" gorm:"primaryKey"`
	// UUID        uuid.UUID `json:"uuid" gorm:"type:uuid;default:gen_random_uuid("`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"createdAt" gorm:"autoCreateTime"`
	UpdatedAt   time.Time `json:"updatedAt" gorm:"autoUpdateTime"`
	ProductID   uint      `json:"productId"`
	Product     Product
}
