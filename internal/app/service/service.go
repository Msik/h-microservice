package service

import (
	"context"

	"github.com/Msik/h-microservice/internal/app/service/category"
	"github.com/Msik/h-microservice/internal/app/service/waste"
	desc "github.com/Msik/h-microservice/pkg/api"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Implementation struct {
	desc.UnimplementedApiServer
	categoryService *category.CategoryService
	wasteService    *waste.WasteService
}

func NewImplementation(categoryService *category.CategoryService, wasteService *waste.WasteService) *Implementation {
	return &Implementation{categoryService: categoryService, wasteService: wasteService}
}

func (impl *Implementation) AddCategoryV1(ctx context.Context, req *desc.AddCategoryV1Request) (*desc.AddCategoryV1Response, error) {
	id, err := impl.categoryService.Add(ctx, req.GetTitle())
	if err != nil {
		return nil, status.Error(codes.Internal, "failed to store catetory")
	}

	return &desc.AddCategoryV1Response{
		Category: &desc.Category{
			Id:    id,
			Title: req.GetTitle(),
		},
	}, nil
}

func (impl *Implementation) CategoryListV1(ctx context.Context, req *desc.CategoryListV1Request) (*desc.CategoryListV1Response, error) {
	panic("unimplemented")
}

func (impl *Implementation) DeleteCategoryV1(ctx context.Context, req *desc.DeleteCategoryV1Request) (*desc.DeleteCategoryV1Response, error) {
	err := impl.categoryService.Delete(ctx, req.GetId())
	if err != nil {
		return nil, status.Error(codes.Internal, "failed to delete catetory")
	}

	return &desc.DeleteCategoryV1Response{
		Success: true,
	}, nil
}

func (impl *Implementation) AddWasteListV1(context.Context, *desc.AddWasteListV1Request) (*desc.AddWasteListV1Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddWasteListV1 not implemented")
}

func (impl *Implementation) DeleteWasteListV1(context.Context, *desc.DeleteWasteListV1Request) (*desc.DeleteWasteListV1Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteWasteListV1 not implemented")
}
