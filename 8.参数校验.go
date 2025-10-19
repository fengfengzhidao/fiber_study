package main

import (
	"fmt"
	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	zh_translations "github.com/go-playground/validator/v10/translations/zh"
	"github.com/gofiber/fiber/v2"
	"strings"
)

var trans ut.Translator
var validate *validator.Validate

func init() {
	// 创建翻译器
	uni := ut.New(zh.New())
	trans, _ = uni.GetTranslator("zh")

	// 初始化 validator 并注册中文翻译
	validate = validator.New()
	if err := zh_translations.RegisterDefaultTranslations(validate, trans); err != nil {
		panic(fmt.Sprintf("注册中文翻译失败：%v", err))
	}
}

func main() {
	app := fiber.New(fiber.Config{})
	app.Get("/", func(c *fiber.Ctx) error {
		type Info struct {
			Name string `query:"name" validate:"required,min=1,max=4"`
			Age  int    `query:"age" validate:"required,min=1,max=100"`
		}
		var info Info
		err := c.QueryParser(&info)
		if err != nil {
			return err
		}
		err = validate.Struct(info)
		if err != nil {
			return err
		}
		return c.JSON(info)
	})

	app.Get("/zh", func(c *fiber.Ctx) error {
		type Info struct {
			Name string `query:"name" validate:"required,min=1,max=4"`
			Age  int    `query:"age" validate:"required,min=1,max=100"`
		}
		var info Info
		err := c.QueryParser(&info)
		if err != nil {
			return err
		}
		err = validate.Struct(info)
		if err != nil {
			errs := err.(validator.ValidationErrors)
			var errMsg []string
			for _, e := range errs {
				errMsg = append(errMsg, e.Translate(trans))
			}
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": strings.Join(errMsg, ";"),
			})
		}
		return c.JSON(info)
	})

	app.Listen(":80")
}
