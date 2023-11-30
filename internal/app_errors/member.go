package app_errors

import "fmt"

type MemberNotFound struct {
	ID    int
	Login string
}

func (u *MemberNotFound) Error() string {
	if u.ID != 0 {
		return fmt.Sprintf("Участник не найден: [ID=%v]", u.ID)
	}
	if u.Login != "" {
		return fmt.Sprintf("Участник не найден: [LOGIN=%v]", u.Login)
	}
	return "Участник не найден"
}

type MemberAlreadyExists struct {
	ID    int
	Login string
}

func (u *MemberAlreadyExists) Error() string {
	if u.ID != 0 {
		return fmt.Sprintf("Участник уже существует: [ID=%v]", u.ID)
	}
	if u.Login != "" {
		return fmt.Sprintf("Участник уже существует: [LOGIN=%v]", u.Login)
	}
	return "Участник уже существует"
}
