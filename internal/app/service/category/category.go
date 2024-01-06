package category

import (
	"context"

	"github.com/jmoiron/sqlx"

	"github.com/Msik/h-microservice/internal/app/model"
	"github.com/Msik/h-microservice/internal/app/repository"
)

type CategoryServiceInterface interface {
	Add(ctx context.Context, title string) (uint64, error)
	List(ctx context.Context) ([]model.Category, error)
	Delete(ctx context.Context, category_id uint64) error
}

type CategoryService struct {
	categoryRepository repository.CategoryRepository
}

func New(db *sqlx.DB) *CategoryService {
	return &CategoryService{categoryRepository: repository.NewCategoryRepository(db)}
}
