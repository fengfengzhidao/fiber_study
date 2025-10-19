package main

import (
	"errors"
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New(fiber.Config{
		CaseSensitive: true, // 路由大小写敏感（如 /User 和 /user 视为不同路由）
		StrictRouting: true, // 严格路由（如 /user/ 和 /user 视为不同路由）
		AppName:       "fengfeng-user-api",
		ErrorHandler: func(c *fiber.Ctx, err error) error {
			return c.Status(200).JSON(map[string]any{
				"code": 1,
				"data": map[string]any{},
				"msg":  err.Error(),
			})
		},
	})
	app.Get("/user", func(c *fiber.Ctx) error {
		return c.SendString("/user-" + c.App().Config().AppName)
	})
	app.Get("/error", func(c *fiber.Ctx) error {
		return errors.New("出错了")
	})
	app.Listen(":80")
}
