package application

import (
	"github.com/gofiber/fiber/v2"
)

// Routes will initialize the routes for the main application
func Routes(app *fiber.App) {
	app.Get("/home", getHome)
	app.Get("", getHome)

	app.Get("/sent", getSent)

	app.Get("/add", getAdd)
}