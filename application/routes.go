package application

import (
	"github.com/gofiber/fiber/v2"
)

// Routes will initialize the routes for the main application
func Routes(app *fiber.App) {
	app.Get("/home", getHome)
	app.Get("", func (c*fiber.Ctx) error {
		return c.Redirect("/home")
	})

	app.Get("/sent", getSent)

	// add kard
	app.Get("/add", getAdd)
	app.Post("/add", postAdd)

	// get kard
	app.Get("/kard/:kardID", getKard)

}