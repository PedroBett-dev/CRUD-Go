package main

import (
	"log"

	"github.com/PedroBett-dev/CRUD-Go.git/src/configuration/database"
	"github.com/PedroBett-dev/CRUD-Go.git/src/controller"
	"github.com/PedroBett-dev/CRUD-Go.git/src/controller/routes"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {

	if err := godotenv.Load(); err != nil {
		log.Println("Warning: .env file not found, using environment variables")
	}

	pool, err := database.NewConnection()
	if err != nil {
		log.Fatalf("Error trying to connect to database: %s", err.Error())
	}
	defer pool.Close()

	controller.InitRepository(pool)

	router := gin.Default()

	routes.InitRoutes(&router.RouterGroup)

	err = router.Run(":2611")
	if err != nil {
		log.Fatal("Inicialization Error: Error to run application")
	}
}