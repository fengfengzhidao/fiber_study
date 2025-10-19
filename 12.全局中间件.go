package main

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"time"
)

func logMiddleware1(c *fiber.Ctx) error {
	t1 := time.Now()
	c.Next()
	subTime := time.Since(t1)
	fmt.Printf("%s %s %s %d %s\n", c.Method(), c.Path(), c.IP(), c.Response().StatusCode(), subTime)
	return nil
}

func homeView1(c *fiber.Ctx) error {
	return c.JSON(map[string]any{
		"path": c.Path(),
		"ip":   c.IP(),
		"ua":   c.Get("User-Agent"),
	})
}

func main() {
	app := fiber.New(fiber.Config{})

	app.Use(logMiddleware1)

	app.Static("/static", "static")

	userGroup := app.Group("/user")
	userGroup.Get("", homeView1)

	app.Listen(":80")
}
