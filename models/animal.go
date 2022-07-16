package models

import (
	"errors"
	"github.com/MeGaNeKoS/TF-Backend/database"
	"strconv"
)

type Animal struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Class string `json:"class"`
	Legs  int    `json:"legs"`
}

type AnimalModel struct{}

func (m *AnimalModel) DeleteAnimal(animal Animal) (err error) {
	db := database.GetSqliteDB()

	result := db.Where("id = ?", animal.ID).Delete(&animal)
	// manual check when trying to delete a non-existing animal
	if result.RowsAffected == 0 {
		return errors.New("No animal with id " + strconv.Itoa(animal.ID))
	}
	// any other error
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (m *AnimalModel) GetAnimalByLimit(offset, limit int) (animals []Animal, err error) {
	db := database.GetSqliteDB()

	result := db.Limit(limit).Offset(offset).Find(&animals)
	if result.Error != nil {
		return []Animal{}, result.Error
	}
	return animals, nil
}

func (m *AnimalModel) GetById(id int) (animal Animal, err error) {
	db := database.GetSqliteDB()
	result := db.Where("id = ?", id).Find(&animal)
	if result.Error != nil {
		return Animal{}, result.Error
	}
	return animal, nil
}


func (m *AnimalModel) InsertAnimal(animal Animal) (insertedAnimal Animal, err error) {
	db := database.GetSqliteDB()
	result := db.Create(&animal)
	if result.Error != nil {
		return Animal{}, result.Error
	}
	return animal, nil
}

func (m *AnimalModel) UpdateAnimal(animal Animal) (updatedAnimal Animal, err error) {
	db := database.GetSqliteDB()

	result := db.Model(&animal).Where("id = ?", animal.ID).Updates(&animal)
	// manual check when trying to update a non-existing animal
	if result.RowsAffected == 0 {
		return Animal{}, errors.New("No animal with id " + strconv.Itoa(animal.ID))
	}
	// any other error
	if result.Error != nil {
		return Animal{}, result.Error
	}
	return animal, nil
}
