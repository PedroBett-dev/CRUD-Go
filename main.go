package main

import (
	"log"

	"github.com/PedroBett-dev/CRUD-Go.git/src/controller/routes"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Load Error: .env file")
	}

	router := gin.Default()

	routes.InitRoutes(&router.RouterGroup)

	err = router.Run(":2611")
	if err != nil {
		log.Fatal("Inicialization Error: Error to run application")
	}
}
