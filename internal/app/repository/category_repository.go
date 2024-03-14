package repository

import (
	"context"

	"github.com/jmoiron/sqlx"

	sq "github.com/Masterminds/squirrel"

	"github.com/Msik/h-microservice/internal/app/model"
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

func (cr *CategoryRepository) GetById(ctx context.Context, id uint64) (model.Category, error) {
	query := sq.Select("id", "title").
		From(catergoryTable).
		Where(sq.Eq{"id": id}).
		RunWith(cr.db).
		PlaceholderFormat(sq.Dollar)

	var category model.Category
	err := query.QueryRowContext(ctx).Scan(&category.Id, &category.Title)
	if err != nil {
		return category, err
	}

	return category, nil
}

func (cr *CategoryRepository) Delete(ctx context.Context, id uint64) error {
	query := sq.Delete(catergoryTable).
		Where(sq.Eq{"id": id}).
		RunWith(cr.db).
		PlaceholderFormat(sq.Dollar)

	var found bool
	err := query.QueryRowContext(ctx).Scan(&found)
	if err != nil {
		return err
	}

	return nil
}
