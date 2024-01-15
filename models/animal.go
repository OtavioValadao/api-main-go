package models

import (
	"time"

	"gorm.io/gorm"
)

type Animal struct {
	gorm.Model
	AnimalGuardianID uint
	Name             string
	Sex              string
	Age              int
	SpeciesAfAniaml  string
	BreedOfAnimal    string
	IsExcluded       bool
}

type AnimalResponse struct {
	ID               uint      `json:"id"`
	CreatedAt        time.Time `json:"created_at"`
	UpdatedAt        time.Time `json:"updated_at"`
	DeletedAt        time.Time `json:"deleted_at,omitempty"`
	AnimalGuardianID uint      `json:"animal_guardian_id"`
	Name             string    `json:"name"`
	Sex              string    `json:"sex"`
	Age              int       `json:"age"`
	SpeciesOfAnimal  string    `json:"species_of_animal"`
	BreedOfAnimal    string    `json:"breed_of_animal"`
	IsExcluded       bool      `json:"is_excluded"`
}
