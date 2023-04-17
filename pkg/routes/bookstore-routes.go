package routes

import (
	"github.com/akshitgautam42/bookstore-go/pkg/controllers"

	"github.com/gofiber/fiber/v2"
)

func SetupBookstoreRoutes(app *fiber.App) {
	// Public routes
	app.Get("/api/books", controllers.GetBooks)
	app.Post("/api/orders", controllers.PlaceOrder)

	// Admin routes
	admin := app.Group("/admin")
	admin.Use(AuthenticateAdmin)
	admin.Get("/api/orders", controllers.GetOrders)
	admin.Post("/api/books", controllers.AddBook)
	admin.Put("/api/books/:id", controllers.UpdateBook)
	admin.Delete("/api/books/:id", controllers.DeleteBook)
}
