package repository

import "go.uber.org/fx"

type exampleParams struct {
	fx.In
}

type exampleRepository struct {
}

func NewExampleRepository(p exampleParams) Repository {
	return &exampleRepository{}
}
