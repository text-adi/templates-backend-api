package service

import (
	"templates/internal/module/entity/repository"

	"go.uber.org/fx"
)

type Params struct {
	fx.In

	Repository repository.Repository
}

type Service struct {
	repo repository.Repository
}

func NewService(p Params) *Service {
	return &Service{
		repo: p.Repository,
	}
}
