package main

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/websocket/v2"
	"io"
)

func main() {
	app := fiber.New()

	app.Get("ws", websocket.New(func(conn *websocket.Conn) {
		addr := conn.RemoteAddr()
		fmt.Println("连接来了", addr)
		for {
			t, byteData, err := conn.ReadMessage()
			if err == io.EOF {
				break
			}
			conn.WriteMessage(t, byteData)
		}
		fmt.Println("连接断开", addr)
	}))

	app.Listen(":80")
}
