package emblem

import (
	"context"
	"guess-emblem/internal/entity"
)

type Service interface {
	GetRandom(ctx context.Context) (entity.Emblem, error)
}

type service struct {
	repo   Repository
}

func (s service) GetRandom(ctx context.Context) (entity.Emblem, error) {
	return s.repo.GetRandom(ctx)
}