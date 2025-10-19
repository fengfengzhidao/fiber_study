package main

import "github.com/gofiber/fiber/v2"

func main() {
	app := fiber.New(fiber.Config{})
	app.Static("/static", "static", fiber.Static{
		Browse: true,
		Index:  "home.html",
	})
	app.Listen(":80")
}
