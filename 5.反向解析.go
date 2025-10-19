package main

import "github.com/gofiber/fiber/v2"

func main() {
	app := fiber.New(fiber.Config{})

	app.Get("/xxx/yyy/zzz/eee/abcdefg/:id", func(c *fiber.Ctx) error { return nil }).Name("u")

	app.Get("/add", func(c *fiber.Ctx) error {
		url, _ := c.GetRouteURL("u", fiber.Map{"id": 1})
		return c.SendString(url)
	})

	app.Listen(":80")
}
