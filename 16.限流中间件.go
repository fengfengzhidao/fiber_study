package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/limiter"
	"time"
)

func GetIp(c *fiber.Ctx) string {
	return c.IP()
}

func main() {
	app := fiber.New()

	// 注册限流中间件
	app.Use(limiter.New(limiter.Config{
		KeyGenerator: GetIp,           // 按 IP 限流（也可按用户 ID 等自定义）
		Expiration:   5 * time.Second, // 限流窗口时间
		Max:          1,               // 窗口内最大请求数
		// 超过限流时的响应
		LimitReached: func(c *fiber.Ctx) error {
			return c.Status(fiber.StatusTooManyRequests).JSON(fiber.Map{
				"code": 429,
				"msg":  "请求过于频繁，请稍后再试",
			})
		},
	}))

	app.Get("/api/limit", func(c *fiber.Ctx) error {
		return c.SendString("限流接口请求成功")
	})
	app.Listen(":80")
}
