package database

import (
	"time"
)

type Client struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"createdAt" gorm:"autoCreateTime"`
	UpdatedAt time.Time `json:"updatedAt" gorm:"autoUpdateTime"`
	Users     []User    `gorm:"many2many:client_users;"`
}

func NewClient() *Client {
	return &Client{}
}

func CreateClient(newUser *Client) error {
	conn, err := GetDbConnection()
	if err != nil {
		return err
	}
	conn.Create(newUser)
	return nil
}

func GetClientByName(name string) (*Client, error) {
	conn, err := GetDbConnection()
	if err != nil {
		return nil, err
	}
	var client = Client{}
	conn.Where("Name = ?", name).Find(&client)
	return &client, nil
}
