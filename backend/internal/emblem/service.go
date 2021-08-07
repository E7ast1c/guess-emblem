package emblem

import (
	"context"
	"guess-emblem/internal/entity"
)

type Service interface {
	GetRandom(ctx context.Context) (entity.Emblem, error)
	Next(ctx context.Context) ([]entity.NextEmblem, error)
}

type service struct {
	repo Repository
}

func NewService(repo Repository) *service {
	return &service{repo: repo}
}

func (s service) GetRandom(ctx context.Context) (entity.Emblem, error) {
	return s.repo.GetRandom(ctx)
}

// Next grab a next emblems for game
func (s service) Next(ctx context.Context) ([]entity.NextEmblem, error) {
	embs, err := s.repo.NextEmblem(ctx, "e2ebee08-e59b-cf5e-e4fd-8f009ccdd231")
	if err != nil {
		return nil, err
	}

	return embs, nil
}
