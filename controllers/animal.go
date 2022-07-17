package controllers

import (
	"github.com/MeGaNeKoS/TF-Backend/models"
	"github.com/gin-gonic/gin"
	"strconv"
)

type AnimalController struct{}

func DeleteAnimal(context *gin.Context) {
	var animal models.Animal
	var err error

	err = context.BindJSON(&animal)
	if err != nil {
		context.JSON(400, gin.H{"status": "error", "message": err})
		context.Abort()
		return
	}

	// delete the animal to the database
	httpCode, err := animalModel.DeleteAnimal(animal)
	if err != nil {
		context.JSON(httpCode, gin.H{"status": "error", "message": err})
		context.Abort()
		return
	}

	context.JSON(httpCode, gin.H{})
}

func GetAnimalById(context *gin.Context) {

	animalID := context.Param("id")
	// convert id to int. if error, return error
	idInt, err := strconv.Atoi(animalID)
	if err != nil {
		context.JSON(400, gin.H{"status": "error", "message": err})
		context.Abort()
		return
	}
	// get the animal from the database
	animal, err := animalModel.GetById(idInt)
	if err != nil {
		context.JSON(400, gin.H{"status": "error", "message": err})
		context.Abort()
		return
	}

	if animal == (models.Animal{}) {
		context.JSON(404, gin.H{"status": "error", "message": "Animal not found"})
		context.Abort()
		return
	}
	context.JSON(200, gin.H{"status": "success", "data": animal})
}

// GetAnimal is for getting list of all animal in database
func GetAnimal(context *gin.Context) {
	var err error

	page := context.Query("page")
	limit := context.Query("limit")

	// convert page and limit to int. if error, set to default value
	pageInt, err := strconv.Atoi(page)
	if err != nil {
		pageInt = 0
	}
	limitInt, err := strconv.Atoi(limit)
	if err != nil {
		limitInt = 10
	}

	animals, err := animalModel.GetAnimalByLimit(pageInt, limitInt)
	if err != nil {
		context.JSON(400, gin.H{"status": "error", "message": err})
		context.Abort()
		return
	}

	if len(animals) == 0 {
		context.JSON(404, gin.H{"status": "error", "message": "Animal not found"})
		context.Abort()
		return

	}

	context.JSON(200, gin.H{"status": "success", "data": animals})
}

func InputAnimal(context *gin.Context) {
	var animal models.Animal
	var err error

	err = context.BindJSON(&animal)
	if err != nil {
		context.JSON(400, gin.H{"status": "error", "message": err})
		context.Abort()
		return
	}

	// insert the animal to the database
	animal, err = animalModel.InsertAnimal(animal)
	if err != nil {
		context.JSON(400, gin.H{"status": "error", "message": "Animal already exists"})
		context.Abort()
		return
	}

	context.JSON(200, gin.H{"status": "success", "data": animal})
}

func UpdateAnimal(context *gin.Context) {
	var animal models.Animal
	var err error

	err = context.BindJSON(&animal)
	if err != nil {
		context.JSON(400, gin.H{"status": "error", "message": err})
		context.Abort()
		return
	}

	// update the animal to the database
	animal, err = animalModel.UpdateAnimal(animal)
	if err != nil {
		context.JSON(400, gin.H{"status": "error", "message": err})
		context.Abort()
		return
	}

	context.JSON(200, gin.H{"status": "success", "data": animal})
}
