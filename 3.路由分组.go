package main

import "github.com/gofiber/fiber/v2"

func main() {
	app := fiber.New(fiber.Config{})
	app.Route("/system", func(router fiber.Router) {
		router.Get("health", func(c *fiber.Ctx) error { return c.SendString("system.health") })
		router.Get("info", func(c *fiber.Ctx) error { return c.SendString("system.info") })
		router.Get("user/list", func(c *fiber.Ctx) error { return c.SendString("system.user/list") })
	})

	group := app.Group("/video")
	group.Get("info", func(c *fiber.Ctx) error { return c.SendString("video.info") })
	group.Get("progress", func(c *fiber.Ctx) error { return c.SendString("video.progress") })

	app.Listen(":80")
}
