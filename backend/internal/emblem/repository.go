package emblem

import (
	"context"
	"guess-emblem/internal/entity"

	"github.com/jmoiron/sqlx"
)

type Repository interface {
	GetRandom(ctx context.Context) (entity.Emblem, error)
	AdditionalName(ctx context.Context, emb entity.Emblem, grain int8) ([]string, error)
}

type repository struct {
	db *sqlx.DB
}

func NewRepository(db *sqlx.DB) *repository {
	return &repository{db: db}
}

func (r repository) GetRandom(ctx context.Context) (entity.Emblem, error) {
	el := entity.Emblem{}
	err := r.db.QueryRow(`select id, name, full_name, title, description, emblem, link 
	from "Emblem" order by random() limit 1`).
		Scan(&el.Id, &el.Name, &el.FullName, &el.Title, &el.Description, &el.Emblem, &el.Link)
	if err != nil {
		return entity.Emblem{}, err
	}
	return el, err
}

// AdditionalName by entity.Emblem return entity.EmblemNext uniq names for game
func (r repository) AdditionalName(ctx context.Context, emb entity.Emblem, grain int8) ([]string, error) {
	names := make([]string, 0, grain)
	err := r.db.Select(&names, `select name from "Emblem" where id != $1 order by random() limit $2`, emb.Id, grain)
	if err != nil {
		return nil, err
	}

	return names, err
}
