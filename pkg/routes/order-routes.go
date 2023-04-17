package routes

import (
	"github.com/akshitgautam42/bookstore-go/pkg/controllers"
	"github.com/gofiber/fiber/v2"
)

func SetupOrderRoutes(app *fiber.App) {
	// Public routes
	app.Get("/api/orders", controllers.GetOrders)

	// Admin routes
	admin := app.Group("/admin")
	admin.Use(AuthenticateAdmin)
	admin.Get("/api/orders/:id", controllers.GetOrder)
	admin.Put("/api/orders/:id", controllers.UpdateOrder)
	admin.Delete("/api/orders/:id", controllers.DeleteOrder)
}
