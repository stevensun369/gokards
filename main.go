package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/django"

	"github.com/stevensun369/kards/database"

	"github.com/stevensun369/kards/conf"


	// routes
	"github.com/stevensun369/kards/auth"
	"github.com/stevensun369/kards/application"


)

func main () {

	engine := django.New(conf.ViewsFolder, ".html")

	app := fiber.New(fiber.Config{
		Views: engine,
	})

	app.Static("/static/", conf.StaticFolder)
	app.Static("/media/", conf.MediaFolder)


	// database initialization
	database.InitDatabase()

	//routes
	auth.Routes(app)
	application.Routes(app)

	app.Listen(":3000")
}