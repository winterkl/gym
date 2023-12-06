package subscription_model

import (
	subscription_entity "awesomeProject/internal/domain/subscription/entity"
)

type CreateSubscriptionDTO struct {
	Title    string `json:"title" binding:"required"`
	Duration int    `json:"duration" binding:"required"`
	Ruble    int    `json:"ruble" binding:"required"`
	Penny    int    `json:"penny"`
}

type GetSubscriptionDTO struct {
	ID       int    `json:"id"`
	Title    string `json:"title"`
	Duration int    `json:"duration"`
	Ruble    int    `json:"ruble"`
	Penny    int    `json:"penny"`
}

func NewCreateMemberDTO(title string, duration, ruble, penny int) CreateSubscriptionDTO {
	return CreateSubscriptionDTO{
		Title:    title,
		Duration: duration,
		Ruble:    ruble,
		Penny:    penny,
	}
}
func NewGetSubscriptionResponse(subscription subscription_entity.Subscription) GetSubscriptionDTO {
	return GetSubscriptionDTO{
		ID:       subscription.ID,
		Title:    subscription.Title,
		Duration: subscription.Duration,
		Ruble:    subscription.Ruble,
		Penny:    subscription.Penny,
	}
}
func NewGetSubscriptionListResponse(subscriptionList []subscription_entity.Subscription) []GetSubscriptionDTO {
	subscriptionListDTO := []GetSubscriptionDTO{}
	for _, subscription := range subscriptionList {
		subscriptionListDTO = append(subscriptionListDTO, NewGetSubscriptionResponse(subscription))
	}
	return subscriptionListDTO
}

type UpdateSubscriptionDTO struct {
	ID       int
	Title    string `json:"title"`
	Duration int    `json:"duration"`
	Ruble    int    `json:"ruble"`
	Penny    int    `json:"penny"`
}
