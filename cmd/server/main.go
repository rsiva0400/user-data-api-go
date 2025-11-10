package main

import (
	"userdata-api/config"
	"userdata-api/internal/handler"
	"userdata-api/internal/logger"
	"userdata-api/internal/repository"
	"userdata-api/internal/routes"
	"userdata-api/internal/service"

	"github.com/gofiber/fiber/v2"
)

func main() {

	// Initialize zap logger
	logger.InitLogger()
	defer logger.Logger.Sync() // flushes buffered logs

	sugar := logger.Sugar()

	// new fiber app
	app := fiber.New()

	// loading configs
	conn := config.InitDB()
	defer conn.Close()

	// set up repository
	userRepository := repository.NewUserRepository(conn)

	// set up service
	userService := service.NewUserService(userRepository)

	// set up userHandler
	userHandler := handler.NewUserHandler(userService)

	// api routes
	routes.SetUpRoutes(app, userHandler)

	// 🔹 Start the server
	sugar.Infow("🚀 Server running on http://localhost:3000")
	app.Listen(":3000")

	sugar.Infow("Starting server...",
		"port", 3000,
		"env", "development",
	)
}
