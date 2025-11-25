package fiber

import (
	"context"
	"fmt"
	"templates/internal/service/fiber/controller"

	"github.com/gofiber/fiber/v2"
	"go.uber.org/fx"
)

type Params struct {
	fx.In

	HttpApp     *fiber.App
	Controllers []controller.Controller `group:"controller"`
}
type Result struct {
	fx.Out

	HttpApp *fiber.App
}

func New() (Result, error) {
	configFiber := fiber.Config{
		StrictRouting:         true,
		DisableStartupMessage: false,
		EnablePrintRoutes:     true,
	}
	return Result{
		HttpApp: fiber.New(configFiber),
	}, nil
}

func Run(lc fx.Lifecycle, p Params) {
	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			for _, r := range p.Controllers {
				r.Register(p.HttpApp)
			}
			go func() {
				if err := p.HttpApp.Listen(fmt.Sprintf(":%d", 3000)); err != nil {
				}
			}()
			return nil
		},
		OnStop: func(ctx context.Context) error {
			return p.HttpApp.Shutdown()
		},
	})

}

var Module = fx.Module(
	"FiberServiceModule",
	fx.Provide(
		New,
	),
	fx.Invoke(
		Run,
	),
)
