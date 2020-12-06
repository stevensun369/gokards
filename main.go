package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/django"

	"github.com/stevensun369/kards/database"

)

func main () {
	engine := django.New("./views", ".html")

	app := fiber.New(fiber.Config{
		Views: engine,
	})

	app.Static("/static/", "./static")

	// database initialization
	database.InitDatabase()

	

}