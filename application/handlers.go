package application

import (
	"github.com/gofiber/fiber/v2"
)

func getHome(c *fiber.Ctx) error {
	return c.Render("application/home", fiber.Map{}, "layouts/main")
}

func getSent(c *fiber.Ctx) error {
	return c.Render("application/sent", fiber.Map{
		"Title": "sent kards",
	}, "layouts/main")
}

func getAdd(c *fiber.Ctx) error {
	return c.Render("application/add", fiber.Map{
		"Title": "create a new kard",
	}, "layouts/main")
}