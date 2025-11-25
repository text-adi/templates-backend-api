package entity

import (
	"templates/internal/module/entity/service"

	"github.com/gofiber/fiber/v2"
)

type Controller struct {
	service service.Service
}

func NewController(s service.Service) *Controller {
	return &Controller{service: s}
}

func (c *Controller) Register(app *fiber.App) {
	app.Get("/", c.GetHelloWorld)
}
