package routes

import (
	"userdata-api/internal/handler"

	"github.com/gofiber/fiber/v2"
)

func SetUpRoutes(api *fiber.App, userHandler *handler.UserHandler) {

	// Group all user-related routes under /users
	user := api.Group("/users")

	// 🧩 User routes
	user.Post("/", userHandler.CreateUser)      // Create new user
	user.Get("/", userHandler.ListUsers)        // Get all users (with pagination)
	user.Get("/:id", userHandler.GetUserByID)   // Get a specific user by ID
	user.Put("/:id", userHandler.UpdateUser)    // Update user
	user.Delete("/:id", userHandler.DeleteUser) // Delete user
}
