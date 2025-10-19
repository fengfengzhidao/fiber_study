package main

import "github.com/gofiber/fiber/v2"

func main() {
	app := fiber.New(fiber.Config{})
	app.Post("/json-form", func(c *fiber.Ctx) error {
		type Info struct {
			Name string `json:"name" form:"name"`
			Age  int    `json:"age" form:"age"`
		}
		var info Info
		err := c.BodyParser(&info)
		if err != nil {
			return err
		}
		return c.JSON(info)
	})
	app.Get("/query", func(c *fiber.Ctx) error {
		type Info struct {
			Name string `query:"name"`
			Age  int    `query:"age"`
		}
		var info Info
		err := c.QueryParser(&info)
		if err != nil {
			return err
		}
		return c.JSON(info)
	})
	app.Get("/user/:name", func(c *fiber.Ctx) error {
		type Info struct {
			Name string `params:"name"`
		}
		var info Info
		err := c.ParamsParser(&info)
		if err != nil {
			return err
		}
		return c.JSON(info)
	})

	app.Listen(":80")
}
