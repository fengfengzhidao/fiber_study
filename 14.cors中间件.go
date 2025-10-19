package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	app := fiber.New(fiber.Config{})

	app.Use(cors.New(cors.Config{
		AllowOrigins:     "https://example.com,http://localhost:63342", // 允许的前端域名
		AllowMethods:     "GET,POST,PUT,DELETE,PATCH",                  // 允许的请求方法
		AllowHeaders:     "Content-Type,Authorization",                 // 允许的请求头
		ExposeHeaders:    "Content-Length",                             // 允许前端读取的响应头
		AllowCredentials: true,                                         // 允许携带 Cookie
	}))

	app.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(map[string]any{"name": "fengfeng"})
	})

	app.Listen(":80")
}
