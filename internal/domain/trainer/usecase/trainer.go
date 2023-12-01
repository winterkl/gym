package trainer_usecase

import (
	"awesomeProject/internal/domain/trainer"
	"awesomeProject/internal/domain/trainer/entity"
	"awesomeProject/internal/domain/trainer/model"
	"context"
	"fmt"
)

type TrainerUseCase struct {
	trainerRepo trainer.Repository
}

func NewTrainerUseCase(trainerRepo trainer.Repository) *TrainerUseCase {
	return &TrainerUseCase{
		trainerRepo: trainerRepo,
	}
}

func (uc *TrainerUseCase) CreateTrainer(ctx context.Context, dto trainer_model.CreateTrainerDTO) error {
	trainer := trainer_entity.NewTrainerFromCreate(dto.MemberID)
	if err := uc.trainerRepo.CreateTrainer(ctx, trainer); err != nil {
		return fmt.Errorf("TrainerUseCase - CreateTrainer - trainerRepo.CreateTrainer: %w", err)
	}
	return nil
}
func (uc *TrainerUseCase) GetTrainer(ctx context.Context, trainerID int) (trainer_model.GetTrainerDTO, error) {
	trainer, err := uc.trainerRepo.GetTrainer(ctx, trainerID)
	if err != nil {
		return trainer_model.GetTrainerDTO{}, fmt.Errorf("TrainerUseCase - GetTrainer - "+
			"trainerRepo.GetTrainer: %w", err)
	}
	return trainer_model.NewGetTrainerResponse(trainer), nil
}
func (uc *TrainerUseCase) GetTrainerList(ctx context.Context) ([]trainer_model.GetTrainerDTO, error) {
	trainer, err := uc.trainerRepo.GetTrainerList(ctx)
	if err != nil {
		return []trainer_model.GetTrainerDTO{}, fmt.Errorf("TrainerUseCase - GetTrainerList - "+
			"trainerRepo.GetTrainerList: %w", err)
	}
	return trainer_model.NewGetTrainerListResponse(trainer), nil
}
func (uc *TrainerUseCase) DeleteTrainer(ctx context.Context, trainerID int) error {
	if err := uc.trainerRepo.DeleteTrainer(ctx, trainerID); err != nil {
		return fmt.Errorf("TrainerUseCase - DeleteTrainer - trainerRepo.DeleteTrainer: %w", err)
	}
	return nil
}
