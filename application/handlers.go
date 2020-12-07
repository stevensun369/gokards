package application

import (
	"github.com/gofiber/fiber/v2"
)

func getHome(c *fiber.Ctx) error {
	return c.Render("application/home", fiber.Map{}, "layouts/main")
}