package main

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New(fiber.Config{})
	app.Post("upload", func(c *fiber.Ctx) error {
		fileHeader, err := c.FormFile("file")
		if err != nil {
			return err
		}
		fmt.Println(fileHeader.Filename)
		return c.SaveFile(fileHeader, "uploads/"+fileHeader.Filename)
	})
	app.Post("uploads", func(c *fiber.Ctx) error {
		form, err := c.MultipartForm()
		if err != nil {
			return err
		}
		for s, headers := range form.File {
			for _, header := range headers {
				fmt.Printf("%s %s\n", s, header.Filename)
			}
		}
		return c.Status(201).SendString("成功")
	})
	app.Listen(":80")
}
