package models

import "gorm.io/gorm"

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
