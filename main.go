package main

import (
	"fmt"
	"github.com/MeGaNeKoS/TF-Backend/database"
	"github.com/MeGaNeKoS/TF-Backend/routes"
	"github.com/gin-gonic/gin"
	"log"
	"os"
)

func main() {
	fmt.Println("Starting")
	// If the file doesn't exist, create it, or append to the file
	file, err := os.OpenFile("logs.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	log.SetOutput(file)

	// initialize db connection
	database.InitSqliteDB()

	route := gin.Default()
	routes.Setup(route)

	err = route.Run()
	if err != nil {
		log.Fatal(err)
	}
}
