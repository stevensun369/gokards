package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/django"

	"github.com/stevensun369/kards/database"

	// routes
	"github.com/stevensun369/kards/auth"
	"github.com/stevensun369/kards/application"


)

func main () {

	engine := django.New("./views", ".html")

	app := fiber.New(fiber.Config{
		Views: engine,
	})

	app.Static("/static/", "./static")
	app.Static("/media/", "./media")


	// database initialization
	database.InitDatabase()

	//routes
	auth.Routes(app)
	application.Routes(app)

	app.Listen(":3000")

}