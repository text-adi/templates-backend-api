package http

import (
	"templates/internal/module/entity"
	"templates/internal/service/fiber"

	"github.com/spf13/cobra"
	"go.uber.org/fx"
)

func Run(cmd *cobra.Command) {
	fx.New(
		fx.Provide(func() *cobra.Command { return cmd }),

		entity.Module,

		fiber.Module,
	).Run()
}
