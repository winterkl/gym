package service_repository

import (
	"awesomeProject/internal/app_errors"
	service_entity "awesomeProject/internal/domain/service/entity"
	"awesomeProject/pkg/postgres"
	"context"
	"database/sql"
	"errors"
	"fmt"
)

type ServiceRepository struct {
	db *postgres.Postgres
}

func NewServiceRepository(db *postgres.Postgres) *ServiceRepository {
	return &ServiceRepository{
		db: db,
	}
}

func (r *ServiceRepository) CreateService(ctx context.Context, service service_entity.Service) error {
	if err := r.db.
		NewInsert().
		Model(&service).
		Scan(ctx); err != nil {
		return fmt.Errorf("ServiceRepository - CreateService - NewInsert: %w", err)
	}
	return nil
}
func (r *ServiceRepository) GetService(ctx context.Context, serviceID int) (service_entity.Service, error) {
	service := service_entity.Service{}
	if err := r.db.
		NewSelect().
		Model(&service).
		Where("id = ?", serviceID).
		Scan(ctx); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return service, &app_errors.ServiceNotFound{ID: serviceID}
		}
		return service, fmt.Errorf("ServiceRepository - GetService - NewSelect: %w", err)
	}
	return service, nil
}
func (r *ServiceRepository) GetServiceList(ctx context.Context) ([]service_entity.Service, error) {
	serviceList := []service_entity.Service{}
	if err := r.db.
		NewSelect().
		Model(&serviceList).
		Scan(ctx); err != nil {
		return []service_entity.Service{}, fmt.Errorf("ServiceRepository - GetServiceList - "+
			"NewSelect: %w", err)
	}
	return serviceList, nil
}
func (r *ServiceRepository) UpdateService(ctx context.Context, service service_entity.Service) error {
	if _, err := r.db.
		NewUpdate().
		Model(&service).
		Set("title = ?", service.Title).
		Set("duration = ?", service.Duration).
		Set("ruble = ?", service.Ruble).
		Set("penny = ?", service.Penny).
		Where("id = ?", service.ID).
		Exec(ctx); err != nil {
		return fmt.Errorf("ServiceRepository - UpdateService - NewUpdate: %w", err)
	}
	return nil
}
func (r *ServiceRepository) DeleteService(ctx context.Context, serviceID int) error {
	if _, err := r.db.
		NewDelete().
		Model((*service_entity.Service)(nil)).
		Where("id = ?", serviceID).
		Exec(ctx); err != nil {
		return fmt.Errorf("ServiceRepository - DeleteService - NewDelete: %w", err)
	}
	return nil
}
