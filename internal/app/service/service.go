package service

import (
	"context"

	"github.com/Msik/h-microservice/internal/app/service/category"

	desc "github.com/Msik/h-microservice/pkg/api"
)

type Implementation struct {
	desc.UnimplementedApiServer
	categoryService *category.CategoryService
}

func NewImplementation(categoryService *category.CategoryService) *Implementation {
	return &Implementation{categoryService: categoryService}
}

func (impl *Implementation) AddCategoryV1(ctx context.Context, req *desc.AddCategoryV1Request) (*desc.AddCategoryV1Response, error) {
	panic("unimplemented")
}

func (impl *Implementation) CategoryListV1(ctx context.Context, req *desc.CategoryListV1Request) (*desc.CategoryListV1Response, error) {
	panic("unimplemented")
}

func (impl *Implementation) DeleteCategoryV1(ctx context.Context, req *desc.DeleteCategoryV1Request) (*desc.DeleteCategoryV1Response, error) {
	panic("unimplemented")
}
