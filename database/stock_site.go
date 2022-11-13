package database

import "time"

// TODO: hello
type StockSite struct {
	ID          uint      `json:"id" gorm:"primaryKey"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"createdAt" gorm:"autoCreateTime"`
	UpdatedAt   time.Time `json:"updatedAt" gorm:"autoUpdateTime"`
	Item        []Item    `gorm:"many2many:stock_site_items"`
	ClientID    uint      `json:"clientId"`
	Client      Client
}
