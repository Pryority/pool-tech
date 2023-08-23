package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	ID       uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primary_key" json:"id,omitempty"`
	Email    string    `gorm:"varchar(255);uniqueIndex;not null" json:"email,omitempty"`
	Password string    `gorm:"not null" json:"password,omitempty"`
}

type CreateUserSchema struct {
	Email    string `json:"email" validate:"required"`
	Password string `json:"password" validate:"required"`
}

// CreateUser creates a new user in the database
func CreateUser(db *gorm.DB, userData CreateUserSchema) (*User, error) {
	newUser := User{
		Email:    userData.Email,
		Password: userData.Password,
	}

	if err := db.Create(&newUser).Error; err != nil {
		return nil, err
	}

	return &newUser, nil
}
