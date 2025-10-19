package main

import "github.com/gofiber/fiber/v2"

func main() {
	app := fiber.New(fiber.Config{})
	app.Static("/xxx", "static")
	app.Get("/string", func(c *fiber.Ctx) error {
		return c.SendString("hello world")
	})
	app.Get("/json1", func(c *fiber.Ctx) error {
		return c.JSON(map[string]any{"code": 0, "msg": "成功"})
	})
	app.Get("/json2", func(c *fiber.Ctx) error {
		type Info struct {
			Name string `json:"name"`
		}
		var info = Info{
			Name: "fengfeng",
		}
		return c.JSON(info)
	})
	app.Get("/html1", func(c *fiber.Ctx) error {
		c.Set("Content-Type", "text/html;charset=utf-8")
		return c.SendString("<h1>你好</h1>")
	})
	app.Get("/html2", func(c *fiber.Ctx) error {
		return c.SendFile("index1.html")
	})
	app.Listen(":80")
}
