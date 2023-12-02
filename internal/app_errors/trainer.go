package app_errors

import "fmt"

type TrainerNotFound struct {
	ID int
}

func (u *TrainerNotFound) Error() string {
	if u.ID != 0 {
		return fmt.Sprintf("Тренер не найден: [ID=%v]", u.ID)
	}
	return "Тренер не найден"
}

type TrainerAlreadyExists struct {
	MemberID int
}

func (u *TrainerAlreadyExists) Error() string {
	if u.MemberID != 0 {
		return fmt.Sprintf("Тренер уже существует: [MEMBERID=%v]", u.MemberID)
	}
	return "Тренер уже существует"
}
