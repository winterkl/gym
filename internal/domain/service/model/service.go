package service_model

import (
	service_entity "awesomeProject/internal/domain/service/entity"
)

type CreateServiceDTO struct {
	Title    string `json:"title" binding:"required"`
	Duration int    `json:"duration" binding:"required"`
	Ruble    int    `json:"ruble" binding:"required"`
	Penny    int    `json:"penny"`
}

type GetServiceDTO struct {
	ID       int    `json:"id"`
	Title    string `json:"title"`
	Duration int    `json:"duration"`
	Ruble    int    `json:"ruble"`
	Penny    int    `json:"penny"`
}

func NewCreateMemberDTO(title string, duration, ruble, penny int) CreateServiceDTO {
	return CreateServiceDTO{
		Title:    title,
		Duration: duration,
		Ruble:    ruble,
		Penny:    penny,
	}
}
func NewGetServiceResponse(service service_entity.Service) GetServiceDTO {
	return GetServiceDTO{
		ID:       service.ID,
		Title:    service.Title,
		Duration: service.Duration,
		Ruble:    service.Ruble,
		Penny:    service.Penny,
	}
}
func NewGetServiceListResponse(serviceList []service_entity.Service) []GetServiceDTO {
	serviceListDTO := []GetServiceDTO{}
	for _, service := range serviceList {
		serviceListDTO = append(serviceListDTO, NewGetServiceResponse(service))
	}
	return serviceListDTO
}

type UpdateServiceDTO struct {
	ID       int
	Title    string `json:"title"`
	Duration int    `json:"duration"`
	Ruble    int    `json:"ruble"`
	Penny    int    `json:"penny"`
}
