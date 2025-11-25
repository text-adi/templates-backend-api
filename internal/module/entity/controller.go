package entity

import "github.com/gofiber/fiber/v2"

type Controller struct {
}

func NewController() *Controller {
	return &Controller{}
}
func (c *Controller) Register(app *fiber.App) {
	app.Get("/", c.GetHelloWorld)
}
