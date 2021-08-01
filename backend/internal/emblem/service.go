package emblem

import (
	"context"
	"guess-emblem/internal/entity"
	"sort"
)

type Service interface {
	GetRandom(ctx context.Context) (entity.Emblem, error)
	Next(ctx context.Context) (entity.EmblemNext, error)
}

type service struct {
	repo Repository
}

func NewService(repo Repository) *service {
	return &service{repo: repo}
}

const nextEmblemGrain int8 = 3


func (s service) GetRandom(ctx context.Context) (entity.Emblem, error) {
	return s.repo.GetRandom(ctx)
}

// Next grab a next emblems for game
func (s service) Next(ctx context.Context) (entity.EmblemNext, error) {
	emb, err := s.repo.GetRandom(ctx)
	if err != nil {
		return entity.EmblemNext{}, err
	}

	names, err := s.repo.AdditionalName(ctx, emb, nextEmblemGrain)
	if err != nil {
		return entity.EmblemNext{}, err
	}
	names = append(names, emb.Name)
	sort.Strings(names)

	return entity.EmblemNext{
		Id:     emb.Id,
		Emblem: emb.Emblem,
		Name:   names,
	}, nil
}
