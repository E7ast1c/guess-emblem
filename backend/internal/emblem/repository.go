package emblem

import (
	"context"
	"guess-emblem/internal/entity"

	"github.com/jmoiron/sqlx"
)

type Repository interface {
	GetRandom(ctx context.Context) (entity.Emblem, error)
	AdditionalName(ctx context.Context, emb entity.Emblem, grain int8) ([]string, error)
	NextEmblem(ctx context.Context, uiid string) ([]entity.NextEmblem, error)
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

// AdditionalName by entity.Emblem return entity.NextEmblem uniq names for game
func (r repository) AdditionalName(ctx context.Context, emb entity.Emblem, count int8) ([]string, error) {
	names := make([]string, 0, count)
	err := r.db.Select(&names, `select name from "Emblem" where id != $1 order by random() limit $2`, emb.Id, count)
	if err != nil {
		return nil, err
	}

	return names, err
}

const limitWrong = 3
func (r repository) NextEmblem(ctx context.Context, uiid string) ([]entity.NextEmblem, error) {
	names := make([]entity.NextEmblem, 0, limitWrong+1)
	err := r.db.Select(&names, `WITH
			t1 (name,emblem) AS (select name, emblem
								from "Emblem" e
										 join "user_score" us on e.id != all (us.shown_emblems)
									AND user_uuid = $1
								order by random()
								limit 1),
			t2 (name,emblem) AS (select name, E''::bytea
									  from "Emblem" e
											  where e.name not like (select name from t1)
									  order by random()
									  limit $2)
			SELECT name, emblem FROM t1 FULL JOIN t2 USING (name, emblem);`, uiid, limitWrong)
	if err != nil {
		return nil, err
	}

	return names, err
}