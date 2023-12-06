package service

import (
	service_entity "awesomeProject/internal/domain/service/entity"
	service_model "awesomeProject/internal/domain/service/model"
	"context"
)

type UseCase interface {
	CreateService(ctx context.Context, dto service_model.CreateServiceDTO) error
	GetService(ctx context.Context, serviceID int) (service_model.GetServiceDTO, error)
	GetServiceList(ctx context.Context) ([]service_model.GetServiceDTO, error)
	UpdateService(ctx context.Context, service service_model.UpdateServiceDTO) error
	DeleteService(ctx context.Context, serviceID int) error
}

type Repository interface {
	CreateService(ctx context.Context, service service_entity.Service) error
	GetService(ctx context.Context, serviceID int) (service_entity.Service, error)
	GetServiceList(ctx context.Context) ([]service_entity.Service, error)
	UpdateService(ctx context.Context, service service_entity.Service) error
	DeleteService(ctx context.Context, serviceID int) error
}
