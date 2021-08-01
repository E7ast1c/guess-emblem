package emblem

import (
	"context"
	"guess-emblem/internal/entity"

	"github.com/jmoiron/sqlx"
)

type Repository interface {
	GetRandom(ctx context.Context) (entity.Emblem, error)
}

type repository struct {
	db *sqlx.DB
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

func NewRepository(db *sqlx.DB) *repository {
	return &repository{db: db}
}
