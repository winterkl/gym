package trainer_model

import "awesomeProject/internal/domain/trainer/entity"

type CreateTrainerDTO struct {
	MemberID int `json:"member_id" binding:"required"`
}

type GetTrainerDTO struct {
	ID  int    `json:"id"`
	FIO string `json:"fio"`
}

func NewGetTrainerResponse(trainer trainer_entity.Trainer) GetTrainerDTO {
	return GetTrainerDTO{
		ID:  trainer.ID,
		FIO: trainer.Member.FIO,
	}
}
func NewGetTrainerListResponse(trainerList []trainer_entity.Trainer) []GetTrainerDTO {
	trainerListDTO := []GetTrainerDTO{}
	for _, trainer := range trainerList {
		trainerListDTO = append(trainerListDTO, NewGetTrainerResponse(trainer))
	}
	return trainerListDTO
}
