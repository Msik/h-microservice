package service

import (
	"github.com/Msik/h-microservice/internal/app/repository"
	desc "github.com/Msik/h-microservice/pkg/api"
)

type Implementation struct {
	desc.UnimplementedApiServer
	factoryRepository *repository.FactoryRepository
}

func NewImplementation(factoryRepository *repository.FactoryRepository) *Implementation {
	return &Implementation{
		factoryRepository: factoryRepository
	}
}

func (impl *Implementation) CategoryListV1(ctx context.Context, in *CategoryListV1Request) (*CategoryListV1Response, error) {

}
