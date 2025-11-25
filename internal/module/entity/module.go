package entity

import (
	httpController "templates/internal/service/fiber/controller"

	"go.uber.org/fx"
)

type Params struct {
	fx.In
}

type Result struct {
	fx.Out

	Controller httpController.Controller `group:"controller"`
}

func New(p Params) (Result, error) {
	return Result{
		Controller: NewController(),
	}, nil
}

var Module = fx.Module("entity",
	fx.Provide(New),
)
