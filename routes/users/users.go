package routes

import (
	"net/http"
	"strconv"
	"web-server/entities"
	"web-server/usecases"

	"github.com/gin-gonic/gin"
)

func SetupRouter(userUC usecases.UserUseCase) *gin.Engine {
	router := gin.Default()

	// GET All Users
	router.GET("/users", func(c *gin.Context) {
		users, err := userUC.GetAll()
		if err != nil {
			c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": "Failed to get users"})
			return
		}
		c.IndentedJSON(http.StatusOK, users)
	})

	// GET User by ID
	router.GET("/users/:id", func(c *gin.Context) {
		id, _ := strconv.Atoi(c.Param("id"))
		user, err := userUC.GetByID(id)
		if err != nil {
			c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": "Failed to get user"})
			return
		}
		if user == nil {
			c.IndentedJSON(http.StatusNotFound, gin.H{"error": "User not found"})
			return
		}
		c.IndentedJSON(http.StatusOK, user)
	})

	// POST Create New User
	router.POST("/users", func(c *gin.Context) {
		var user entities.User

		if err := c.ShouldBindJSON(&user); err != nil {
			c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "Invalid user data"})
			return
		}
		if err := userUC.CreateNew(&user); err != nil {
			c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user"})
			return
		}
		c.Status(http.StatusCreated)
	})

	// DELETE User by ID
	router.DELETE("/users/:id", func(c *gin.Context) {
		id, _ := strconv.Atoi(c.Param("id"))
		if err := userUC.DeleteByID(id); err != nil {
			c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete user"})
			return
		}
		c.Status(http.StatusNoContent)
	})

	return router
}
