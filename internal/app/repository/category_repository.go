package repository

import (
	"github.com/Msik/h-microservice/internal/app/model"
	"github.com/jmoiron/sqlx"
)

type CategoryRepository struct {
	db *sqlx.DB
}

func newCategoryRepository(db *sqlx.DB) *CategoryRepository {
	return &CategoryRepository{
		db: db
	}
}

func GetCategory() model.Category {

}
