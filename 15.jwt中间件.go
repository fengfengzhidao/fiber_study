package main

import (
	"github.com/gofiber/fiber/v2"
	jwtware "github.com/gofiber/jwt/v3"
	"github.com/golang-jwt/jwt/v4"
	"time"
)

func main() {
	app := fiber.New(fiber.Config{})

	// 公开接口（无需认证）：登录生成 Token
	app.Post("/login", func(c *fiber.Ctx) error {
		// 模拟登录成功，生成 JWT Token
		token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
			"username": "test",
			"exp":      jwt.NewNumericDate(time.Now().Add(24 * time.Hour)), // 24小时过期
		})
		tokenString, _ := token.SignedString([]byte("your-secret-key")) // 密钥（生产环境需保密）
		return c.JSON(fiber.Map{"token": tokenString})
	})

	// 受保护接口：注册 JWT 中间件
	api := app.Group("").Use(jwtware.New(jwtware.Config{
		SigningKey: []byte("your-secret-key"),
	}))

	api.Get("user", func(c *fiber.Ctx) error {
		return c.SendString("user")
	})

	app.Listen(":80")
}
