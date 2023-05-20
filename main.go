package main

import (
	// route handler
	"web-server/routes"

	// The Gin web framework
	"github.com/gin-gonic/gin"
)

func main() {
	// initialise gin Engine
	router := gin.Default() // a gin router to handle requests

	// request handlers
	router.GET("/people", routes.GetPeople)
	router.POST("/people", routes.PostPeople)
	router.GET("/people/:id", routes.GetPersonByID)

	// Listen at http://localhost:8080
	router.Run(":8080")
}
