package subscription

import (
	subscription_entity "awesomeProject/internal/domain/subscription/entity"
	subscription_model "awesomeProject/internal/domain/subscription/model"
	"context"
)

type UseCase interface {
	CreateSubscription(ctx context.Context, dto subscription_model.CreateSubscriptionDTO) error
	GetSubscription(ctx context.Context, subscriptionID int) (subscription_model.GetSubscriptionDTO, error)
	GetSubscriptionList(ctx context.Context) ([]subscription_model.GetSubscriptionDTO, error)
	UpdateSubscription(ctx context.Context, subscription subscription_model.UpdateSubscriptionDTO) error
	DeleteSubscription(ctx context.Context, subscriptionID int) error
}

type Repository interface {
	CreateSubscription(ctx context.Context, subscription subscription_entity.Subscription) error
	GetSubscription(ctx context.Context, subscriptionID int) (subscription_entity.Subscription, error)
	GetSubscriptionList(ctx context.Context) ([]subscription_entity.Subscription, error)
	UpdateSubscription(ctx context.Context, subscription subscription_entity.Subscription) error
	DeleteSubscription(ctx context.Context, subscriptionID int) error
}
