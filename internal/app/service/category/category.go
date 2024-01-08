package category

import (
	"context"

	"github.com/jmoiron/sqlx"

	"github.com/Msik/h-microservice/internal/app/model"
	"github.com/Msik/h-microservice/internal/app/repository"
)

type CategoryService struct {
	categoryRepository *repository.CategoryRepository
}

func New(db *sqlx.DB) *CategoryService {
	return &CategoryService{categoryRepository: repository.NewCategoryRepository(db)}
}

func (cs *CategoryService) Add(ctx context.Context, title string) (uint64, error) {
	id, err := cs.categoryRepository.Store(ctx, title)
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (cs *CategoryService) List(ctx context.Context) ([]model.Category, error) {
	return []model.Category{}, nil
}

func (cs *CategoryService) Delete(ctx context.Context, category_id uint64) error {
	if err := cs.categoryRepository.Delete(ctx, category_id); err != nil {
		return err
	}

	return nil
}
