package trainer

import (
	"awesomeProject/internal/domain/trainer/entity"
	"awesomeProject/internal/domain/trainer/model"
	"context"
)

type UseCase interface {
	CreateTrainer(ctx context.Context, dto trainer_model.CreateTrainerDTO) error
	GetTrainer(ctx context.Context, trainerID int) (trainer_model.GetTrainerDTO, error)
	GetTrainerList(ctx context.Context) ([]trainer_model.GetTrainerDTO, error)
	DeleteTrainer(ctx context.Context, trainerID int) error
}

type Repository interface {
	CreateTrainer(ctx context.Context, trainer trainer_entity.Trainer) error
	GetTrainer(ctx context.Context, trainerID int) (trainer_entity.Trainer, error)
	GetTrainerList(ctx context.Context) ([]trainer_entity.Trainer, error)
	DeleteTrainer(ctx context.Context, trainerID int) error
}
