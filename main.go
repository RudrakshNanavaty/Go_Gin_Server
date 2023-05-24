package main

import (
	userDatabase "web-server/database/users"
	userRouter "web-server/routes/users"
	userUseCase "web-server/usecases/user"
)

func main() {
	// Initialize repository
	userDB := userDatabase.NewUserDB()

	// Initialize use case
	userUC := userUseCase.NewUserUseCase(userDB)

	// Set up router
	router := userRouter.SetupRouter(userUC)

	// Run the server
	router.Run(":8080")
}
