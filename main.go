package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/django"

	"github.com/stevensun369/kards/database"

	// routes
	"github.com/stevensun369/kards/auth"

)

func main () {

	engine := django.New("./views", ".html")

	app := fiber.New(fiber.Config{
		Views: engine,
	})

	app.Static("/static/", "./static")

	// database initialization
	database.InitDatabase()

	app.Get("/test", func (c * fiber.Ctx) error {
		return c.SendString("Hello, world")
	})

	//routes
	auth.Routes(app)

	app.Listen(":3000")

}