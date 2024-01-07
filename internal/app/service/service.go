package service

import (
	"context"

	"github.com/Msik/h-microservice/internal/app/service/category"
	desc "github.com/Msik/h-microservice/pkg/api"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Implementation struct {
	desc.UnimplementedApiServer
	categoryService *category.CategoryService
}

func NewImplementation(categoryService *category.CategoryService) *Implementation {
	return &Implementation{categoryService: categoryService}
}

func (impl *Implementation) AddCategoryV1(ctx context.Context, req *desc.AddCategoryV1Request) (*desc.AddCategoryV1Response, error) {
	title := reg.GetTitle()
	if reg.GetTitle() == "" {
		return nil, status.Error(codes.InvalidArgument, "title is required")
	}

	id, err := impl.categoryService.Add(ctx, title)
	if err != nil {
		return nil, status.Error(codes.Internal, "failed to store catetory")
	}

	return &desc.AddCategoryV1Response{
		Category: &desc.Category{
			Id: id,
			Title: title,
		}
	}, nil
}

func (impl *Implementation) CategoryListV1(ctx context.Context, req *desc.CategoryListV1Request) (*desc.CategoryListV1Response, error) {
	panic("unimplemented")
}

func (impl *Implementation) DeleteCategoryV1(ctx context.Context, req *desc.DeleteCategoryV1Request) (*desc.DeleteCategoryV1Response, error) {
	if req.GetId() == 0 {
		return nil, status.Error(codes.InvalidArgument, "validation error")
	}

	err := impl.categoryService.Delete(ctx, req.GetId())
	if err != nil {
		return nil, status.Error(codes.Internal, "failed to delete catetory")
	}

	return &desc.DeleteCategoryV1Response{
		Success: true,
	}
}
