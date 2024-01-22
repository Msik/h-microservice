package repository

import (
	"context"

	sq "github.com/Masterminds/squirrel"
	"github.com/jmoiron/sqlx"
)

type WasteRepository struct {
	db *sqlx.DB
}

func NewWasteRepository(db *sqlx.DB) *WasteRepository {
	return &WasteRepository{
		db: db,
	}
}

var (
	wasteTable = "wastes"
)

func (wr *WasteRepository) Store(ctx context.Context, amount float32, category_id uint64) (uint64, error) {
	query := sq.Insert(wasteTable).
		Columns("amount", "category_id").
		Suffix("RETURNING \"id\"").
		Values(amount, category_id).
		RunWith(wr.db).
		PlaceholderFormat(sq.Dollar)

	var id uint64
	err := query.QueryRowContext(ctx).Scan(&id)
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (wr *WasteRepository) Delete(ctx context.Context, id uint64) error {
	query := sq.Delete(wasteTable).
		Where(sq.Eq{"id": id}).
		RunWith(wr.db).
		PlaceholderFormat(sq.Dollar)

	var found bool
	err := query.QueryRowContext(ctx).Scan(&found)
	if err != nil {
		return err
	}

	return nil
}
