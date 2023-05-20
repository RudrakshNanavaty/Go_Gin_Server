package routes

import (
	"fmt"

	// for http statuses
	"net/http"

	// The Gin web framework
	"github.com/gin-gonic/gin"

	// contains the person struct
	"web-server/entities"

	// contains the in-memory DB
	"web-server/database"
)

// respond with the entire database.People struct as JSON
func GetPeople(context *gin.Context) {
	// IndentedJSON makes it look better
	context.IndentedJSON(http.StatusOK, database.People)
}

// add to database.People from JSON received in the request body
func PostPeople(context *gin.Context) {
	var newPerson entities.Person

	// BindJSON to bind the received JSON to newPerson
	if err := context.BindJSON(&newPerson); err != nil {
		// log the error, respond and return
		fmt.Println(err)

		context.IndentedJSON(http.StatusBadRequest, gin.H{
			"message": "Invalid request",
		})

		return
	}

	// append the new entities.Person to database.People
	database.People = append(database.People, newPerson)

	// respond as IndentedJSON
	context.IndentedJSON(http.StatusCreated, newPerson)
}

// locate the entities.Person whose ID value matches the id sent
// then return that entities.Person as a response
func GetPersonByID(context *gin.Context) {
	// get the id from request params
	var id string = context.Param("id")

	// Linear Search through database.People
	for _, p := range database.People {
		// respond and return if ID matched
		if p.ID == id {
			context.IndentedJSON(http.StatusOK, p)
			return
		}
	}

	// respond 404
	context.IndentedJSON(
		http.StatusNotFound,

		// refer https://pkg.go.dev/github.com/gin-gonic/gin#H
		gin.H{
			"message": "entities.Person not found",
		},
	)
}
