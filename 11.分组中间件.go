package main

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"time"
)

func logMiddleware(c *fiber.Ctx) error {
	t1 := time.Now()
	c.Next()
	subTime := time.Since(t1)
	fmt.Printf("%s %s %s %d %s\n", c.Method(), c.Path(), c.IP(), c.Response().StatusCode(), subTime)
	return nil
}

func homeView(c *fiber.Ctx) error {
	return c.JSON(map[string]any{
		"path": c.Path(),
		"ip":   c.IP(),
		"ua":   c.Get("User-Agent"),
	})
}

func main() {
	app := fiber.New(fiber.Config{})

	userGroup := app.Group("/user", logMiddleware)
	userGroup.Get("", homeView)
	userGroup.Get("info", homeView)

	app.Route("/video", func(router fiber.Router) {
		router.Use(logMiddleware)
		router.Get("", homeView)
		router.Get("info", homeView)
	})

	app.Get("no", homeView)

	app.Listen(":80")
}
