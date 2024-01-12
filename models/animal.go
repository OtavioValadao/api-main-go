package models

import "gorm.io/gorm"

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
