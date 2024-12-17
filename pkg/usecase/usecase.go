package usecase

import "testing-system-api/pkg/service"

type Usecase struct {
	services *service.Service
}

func NewUsecase(service *service.Service) *Usecase {
	return &Usecase{services: service}
}
