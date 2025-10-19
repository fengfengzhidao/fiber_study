package main

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

var validate1 *validator.Validate

func init() {
	validate1 = validator.New()
}
func bindMiddleware[T any](c *fiber.Ctx) error {
	var cr T
	err := c.QueryParser(&cr)
	if err != nil {
		return err
	}
	err = validate1.Struct(cr)
	if err != nil {
		return err
	}
	c.Locals("value", cr)
	c.Next()
	return nil
}

func main() {
	app := fiber.New(fiber.Config{})

	type Info struct {
		Name string `query:"name"`
	}

	app.Get("/", bindMiddleware[Info], func(c *fiber.Ctx) error {
		info := c.Locals("value").(Info)
		return c.JSON(info)
	})

	app.Listen(":80")
}
