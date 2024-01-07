package repository

import (
	"context"

	"github.com/jmoiron/sqlx"

	sq "github.com/Masterminds/squirrel"
)

type CategoryRepository struct {
	db *sqlx.DB
}

func NewCategoryRepository(db *sqlx.DB) *CategoryRepository {
	return &CategoryRepository{
		db: db,
	}
}

var (
	catergoryTable = "categories"
)

func (cr *CategoryRepository) Store(ctx context.Context, title string) (uint64, error) {
	query := sq.Insert(catergoryTable).
		Columns("title").
		Suffix("RETURNING \"id\"").
		Values(title).
		RunWith(cr.db).
		PlaceholderFormat(sq.Dollar)

	var id uint64
	err := query.QueryRowContext(ctx).Scan(&id)
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (cr *CategoryRepository) Delete(ctx context.Context, id uint64) error {
	query := sq.Delete(catergoryTable).
		Where(sq.Eq{"id": id}).
		RunWith(cr.db).
		PlaceholderFormat(sq.Dollar)

	err := query.QueryRowContext(ctx)
	if err != nil {
		return err
	}

	return nil
}
