package service

import (
	desc "github.com/Msik/h-microservice/pkg/api"
)

type Implementation struct {
	desc.UnimplementedApiServer
}

func NewImplementation() *Implementation {
	return &Implementation{}
}
