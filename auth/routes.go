package auth

import (
	"github.com/gofiber/fiber/v2"
)

// Routes will initialize the auth routes
func Routes(app * fiber.App) {
	g := app.Group("auth")

	g.Get("/register", getRegister)
	g.Post("/register", postRegister)
	
	g.Get("/login", getLogin)
	g.Post("/login", postLogin)

	g.Get("/logout", getLogout)


}

