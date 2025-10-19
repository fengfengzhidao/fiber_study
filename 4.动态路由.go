package main

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New(fiber.Config{})
	app.Get("/user/:id", func(c *fiber.Ctx) error {
		id := c.Params("id")
		fmt.Println(id)
		return c.SendString("user_id = " + id)
	})
	app.Get("/user/:id/:name", func(c *fiber.Ctx) error {
		id := c.Params("id")
		name := c.Params("name")
		fmt.Println(id, name)
		return c.SendString(fmt.Sprintf("user_id = %s name = %s", id, name))
	})
	app.Get("*", func(c *fiber.Ctx) error {
		return c.SendString("任意get的路由")
	})
	app.Listen(":80")
}
