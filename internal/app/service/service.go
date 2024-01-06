package service

import (
	"context"

	desc "github.com/Msik/h-microservice/pkg/api"
	"github.com/Msik/h-microservice/internal/app/service/category"
)

type Implementation struct {
	desc.UnimplementedApiServer
	categoryService: category.CategoryServiceInterface
}

func NewImplementation(categoryService CategoryServiceInterface) *Implementation {
	return &Implementation{categoryService: categoryService}
}

func (impl *Implementation) AddCategoryV1(ctx context.Context, req *AddCategoryV1Request) (*AddCategoryV1Response, error) {
	panic("unimplemented")
}

func (impl *Implementation) CategoryListV1(ctx context.Context, req *CategoryListV1Request) (*CategoryListV1Response, error) {
	panic("unimplemented")
}

func (impl *Implementation) DeleteCategoryV1(ctx context.Context, req *DeleteCategoryV1Request) (*DeleteCategoryV1Response, error) {
	panic("unimplemented")
}
