package trainer_repository

import (
	"awesomeProject/internal/app_errors"
	trainer_entity "awesomeProject/internal/domain/trainer/entity"
	"awesomeProject/pkg/postgres"
	"context"
	"database/sql"
	"errors"
	"fmt"
	"github.com/jackc/pgconn"
)

type TrainerRepository struct {
	db *postgres.Postgres
}

func NewTrainerRepository(db *postgres.Postgres) *TrainerRepository {
	return &TrainerRepository{
		db: db,
	}
}

func (r *TrainerRepository) CreateTrainer(ctx context.Context, trainer trainer_entity.Trainer) error {
	if err := r.db.NewInsert().Model(&trainer).Scan(ctx); err != nil {
		if pgErr := err.(*pgconn.PgError); pgErr.Code == r.db.Errors.CodeUniqueConstraint {
			return &app_errors.TrainerAlreadyExists{MemberID: trainer.MemberID}
		}
		return fmt.Errorf("TrainerRepository - CreateTrainer - NewInsert: %w", err)
	}
	return nil
}
func (r *TrainerRepository) GetTrainer(ctx context.Context, trainerID int) (trainer_entity.Trainer, error) {
	var trainer trainer_entity.Trainer
	if err := r.db.NewSelect().Model(&trainer).Relation("Member").
		Where("trainer.id = ?", trainerID).Scan(ctx); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return trainer, &app_errors.TrainerNotFound{ID: trainerID}
		}
		return trainer, fmt.Errorf("TrainerRepository - GetTrainer - NewSelect: %w", err)
	}
	return trainer, nil
}
func (r *TrainerRepository) GetTrainerList(ctx context.Context) ([]trainer_entity.Trainer, error) {
	trainerList := []trainer_entity.Trainer{}
	if err := r.db.NewSelect().Model(&trainerList).Relation("Member").Scan(ctx); err != nil {
		return []trainer_entity.Trainer{}, fmt.Errorf("TrainerRepository - GetTrainerList - NewSelect: %w", err)
	}
	return trainerList, nil
}
func (r *TrainerRepository) DeleteTrainer(ctx context.Context, trainerID int) error {
	if _, err := r.db.NewDelete().Model((*trainer_entity.Trainer)(nil)).
		Where("id = ?", trainerID).Exec(ctx); err != nil {
		return fmt.Errorf("TrainerRepository - DeleteTrainer - NewDelete: %w", err)
	}
	return nil
}
