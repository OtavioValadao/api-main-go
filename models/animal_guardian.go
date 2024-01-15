package models

import (
	"time"

	"gorm.io/gorm"
)

type AnimalGuardian struct {
	gorm.Model
	Animals       []Animal
	Name          string
	Cpf           int
	ContactNumber int
	Email         string
	Address       string
	HouseNumber   int
	Complement    string
}

type AnimalGuardianResponse struct {
	ID            uint             `json:"id"`
	CreatedAt     time.Time        `json:"createdAt"`
	UpdatedAt     time.Time        `json:"updatedAt"`
	DeletedAt     time.Time        `json:"deletedAt,omitempty"`
	Name          string           `json:"name"`
	Cpf           int              `json:"cpf"`
	ContactNumber int              `json:"contact_number"`
	Email         string           `json:"email"`
	Address       string           `json:"address"`
	HouseNumber   int              `json:"house_number"`
	Complement    string           `json:"complement"`
	Animals       []AnimalResponse `json:"animals"`
}
