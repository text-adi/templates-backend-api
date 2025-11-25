package entity

import "github.com/gofiber/fiber/v2"

func (c *Controller) GetHelloWorld(ctx *fiber.Ctx) error {
	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{"Hello": "World"})
}
