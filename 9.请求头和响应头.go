package main

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"net/url"
)

func main() {
	app := fiber.New(fiber.Config{})

	app.Get("/", func(c *fiber.Ctx) error {
		fmt.Println(c.Get("User-Agent"))
		c.Set("X-App", "fengfeng")
		c.Set("X-AppName", url.QueryEscape("枫枫"))
		// fmt.Println(url.QueryUnescape("%E6%9E%AB%E6%9E%AB"))
		return c.SendString("hello")
	})

	app.Listen(":80")
}
