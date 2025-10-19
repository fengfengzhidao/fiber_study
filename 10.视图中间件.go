package main

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
)

func authMiddleware1(c *fiber.Ctx) error {
	fmt.Println("请求来了1")
	//c.SendString("被拦截了")
	//return nil
	fmt.Println("请求体", string(c.Body()))
	c.Next()
	fmt.Println("响应来了1")
	fmt.Println("响应体", string(c.Response().Body()))
	return nil
}
func authMiddleware2(c *fiber.Ctx) error {
	fmt.Println("请求来了2")
	c.Next()
	fmt.Println("响应来了2")
	return nil
}

func indexView(c *fiber.Ctx) error {
	fmt.Println("视图/")
	return c.JSON(map[string]any{"code": 0})
}

func main() {
	app := fiber.New(fiber.Config{})

	app.Get("/", authMiddleware1, authMiddleware2, indexView)

	app.Listen(":80")
}
