package models

import "time"

type User struct {
	ID        uint `gorm:"primary_key" json:"id"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time `sql:"index"`

	Name              string `json:"name"`
	Email             string `json:"email"`
	Mobile            string `json:"mobile"`
	EncryptedPassword string `json:"encrypted_password"`
}
