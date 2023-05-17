package main

import (
	"fmt"

	// for http statuses
	"net/http"

	// The Gin web framework
	"github.com/gin-gonic/gin"
)

// data definitions
type person struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

// people slice to seed record person data.
var people = []person{
	{
		ID:   "1",
		Name: "ABC",
	},
	{
		ID:   "2",
		Name: "DEF",
	},
	{
		ID:   "3",
		Name: "GHI",
	},
}

func main() {
	// initialise gin Engine
	router := gin.Default() // a gin router to handle requests

	// request handlers
	router.GET("/people", getPeople)
	router.POST("/people", postPeople)
	router.GET("/people/:id", getPersonByID)

	// Listen at http://localhost:8080
	router.Run(":8080")
}

// respond with the entire people struct as JSON
func getPeople(context *gin.Context) {
	// IndentedJSON makes it look better
	context.IndentedJSON(http.StatusOK, people)
}

// add to people from JSON received in the request body
func postPeople(context *gin.Context) {
	var newPerson person

	// BindJSON to bind the received JSON to newPerson
	if err := context.BindJSON(&newPerson); err != nil {
		// log the error, respond and return
		fmt.Println(err)

		context.IndentedJSON(http.StatusBadRequest, gin.H{
			"message": "Invalid request",
		})

		return
	}

	// append the new person to people
	people = append(people, newPerson)

	// respond as IndentedJSON
	context.IndentedJSON(http.StatusCreated, newPerson)
}

// locate the person whose ID value matches the id sent
// then return that person as a response
func getPersonByID(context *gin.Context) {
	// get the id from request params
	var id string = context.Param("id")

	// Linear Search through people
	for _, p := range people {
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
			"message": "person not found",
		},
	)
}
