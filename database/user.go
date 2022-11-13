package database

import (
	"time"
)

type User struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	Username  string    `json:"username"`
	Email     string    `json:"email"`
	Password  string    `json:"password"`
	CreatedAt time.Time `json:"createdAt" gorm:"autoCreateTime"`
	UpdatedAt time.Time `json:"updatedAt" gorm:"autoUpdateTime"`
}

func NewUser() *User {
	return &User{}
}

func CreateUser(newUser *User) error {
	conn, err := GetDbConnection()
	if err != nil {
		return err
	}
	conn.Create(newUser)
	return nil
}

func GetUserByUsername(username string) (*User, error) {
	conn, err := GetDbConnection()
	if err != nil {
		return nil, err
	}
	var user = User{}
	conn.Where("Username = ?", username).Find(&user)
	return &user, nil
}
