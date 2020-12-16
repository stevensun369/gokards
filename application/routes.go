package application

import (
	"github.com/gofiber/fiber/v2"
)

// Routes will initialize the routes for the main application
func Routes(app *fiber.App) {
	// register, login and logout redirects
	app.Get("/register", func (c *fiber.Ctx) error {
		return c.Redirect("/auth/register")
	})

	app.Get("/login", func (c *fiber.Ctx) error {
		return c.Redirect("/auth/login")
	})

	app.Get("/logout", func (c *fiber.Ctx) error {
		return c.Redirect("/auth/logout")
	})

	// normal for application
	app.Get("/home", getHome)
	app.Get("", func (c *fiber.Ctx) error {
		return c.Redirect("/home")
	})

	app.Get("/sent", getSent)

	// add kard
	app.Get("/add", getAdd)
	app.Post("/add", postAdd)

	// get kard
	app.Get("/kard/:kardID", getKard)

	// help
	app.Get("/help", getHelp)

}