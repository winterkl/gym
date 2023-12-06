package service_usecase

import (
	"awesomeProject/internal/domain/service"
	service_entity "awesomeProject/internal/domain/service/entity"
	service_model "awesomeProject/internal/domain/service/model"
	"context"
	"fmt"
)

type ServiceUseCase struct {
	serviceRepo service.Repository
}

func NewServiceUseCase(serviceRepo service.Repository) *ServiceUseCase {
	return &ServiceUseCase{serviceRepo: serviceRepo}
}

func (uc *ServiceUseCase) CreateService(ctx context.Context, dto service_model.CreateServiceDTO) error {
	service := service_entity.NewServiceFromCreate(dto.Title, dto.Duration, dto.Ruble, dto.Penny)
	if err := uc.serviceRepo.CreateService(ctx, service); err != nil {
		return fmt.Errorf("ServiceUseCase - CreateService - "+
			"serviceRepo.CreateService: %w", err)
	}
	return nil
}
func (uc *ServiceUseCase) GetService(ctx context.Context, serviceID int) (service_model.GetServiceDTO, error) {
	service, err := uc.serviceRepo.GetService(ctx, serviceID)
	if err != nil {
		return service_model.GetServiceDTO{}, fmt.Errorf("ServiceUseCase - GetService - "+
			"serviceRepo.GetService: %w", err)
	}
	return service_model.NewGetServiceResponse(service), nil
}
func (uc *ServiceUseCase) GetServiceList(ctx context.Context) ([]service_model.GetServiceDTO, error) {
	serviceList, err := uc.serviceRepo.GetServiceList(ctx)
	if err != nil {
		return []service_model.GetServiceDTO{}, fmt.Errorf("ServiceUseCase - "+
			"GetServiceList - serviceRepo.GetServiceList: %w", err)
	}
	return service_model.NewGetServiceListResponse(serviceList), nil
}
func (uc *ServiceUseCase) UpdateService(ctx context.Context, dto service_model.UpdateServiceDTO) error {
	service := service_entity.NewServiceFromUpdate(dto.Title, dto.ID, dto.Duration, dto.Ruble, dto.Penny)
	if err := uc.serviceRepo.UpdateService(ctx, service); err != nil {
		return fmt.Errorf("ServiceUseCase - UpdateService - serviceRepo.UpdateService: %w", err)
	}
	return nil
}
func (uc *ServiceUseCase) DeleteService(ctx context.Context, serviceID int) error {
	if err := uc.serviceRepo.DeleteService(ctx, serviceID); err != nil {
		return fmt.Errorf("ServiceUseCase - DeleteService - serviceRepo.DeleteService: %w", err)
	}
	return nil
}
