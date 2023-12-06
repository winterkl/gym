package app_errors

import "fmt"

type ServiceNotFound struct {
	ID    int
	Title string
}

func (u *ServiceNotFound) Error() string {
	if u.ID != 0 {
		return fmt.Sprintf("Услуга не найден: [ID=%v]", u.ID)
	}
	if u.Title != "" {
		return fmt.Sprintf("Услуга не найден: [TITLE=%v]", u.Title)
	}
	return "Услуга не найден"
}

type ServiceAlreadyExists struct {
	ID    int
	Title string
}

func (u *ServiceAlreadyExists) Error() string {
	if u.ID != 0 {
		return fmt.Sprintf("Услуга уже существует: [ID=%v]", u.ID)
	}
	if u.Title != "" {
		return fmt.Sprintf("Услуга уже существует: [TITLE=%v]", u.Title)
	}
	return "Услуга уже существует"
}
