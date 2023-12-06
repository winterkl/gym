package app_errors

import "fmt"

type SubscriptionNotFound struct {
	ID    int
	Title string
}

func (u *SubscriptionNotFound) Error() string {
	if u.ID != 0 {
		return fmt.Sprintf("Абонемент не найден: [ID=%v]", u.ID)
	}
	if u.Title != "" {
		return fmt.Sprintf("Абонемент не найден: [TITLE=%v]", u.Title)
	}
	return "Абонемент не найден"
}

type SubscriptionAlreadyExists struct {
	ID    int
	Title string
}

func (u *SubscriptionAlreadyExists) Error() string {
	if u.ID != 0 {
		return fmt.Sprintf("Абонемент уже существует: [ID=%v]", u.ID)
	}
	if u.Title != "" {
		return fmt.Sprintf("Абонемент уже существует: [TITLE=%v]", u.Title)
	}
	return "Абонемент уже существует"
}
