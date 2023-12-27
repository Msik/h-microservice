package repository

import (
	"github.com/jmoiron/sqlx"
)

type FactoryRepository struct {
	db *sqlx.DB
}

func NewFactory(db *sqlx.DB) *FactoryRepository {
	return &FactoryRepository{
		db: db
	}
}

func (fr *FactoryRepository) GetCategoryRepository() *CategoryRepository {
	return newCategoryRepository(fr.db)
}
