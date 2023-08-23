package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/pryority/pool-tech/cmd/server/handlers"
)

const LocalPort = ":8080"

// SetupRoutes defines all routes and their corresponding handlers.
func SetupRoutes(app *fiber.App) {
	// Serve static files
	app.Static("/public", "./cmd/web/public")

	// Group related routes with subrouter
	api := app.Group("/api")

	// Define routes
	app.Get("/", handlers.LoadMain)
	// app.Get("/debug", handlers.LoadDebug)

	// Handle API routes
	api.Post("/register", handlers.RegisterUser)
	api.Post("/login", handlers.Login)
	// api.Get("/notes", handlers.GetAllNotes)

	// Catch-all route for templates
	app.Get("/:route_name", handlers.LoadTemplate)
	// app.Get("/refresh/notes", refreshNotes)
}
