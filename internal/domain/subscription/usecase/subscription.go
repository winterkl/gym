package subscription_usecase

import (
	"awesomeProject/internal/domain/subscription"
	subscription_entity "awesomeProject/internal/domain/subscription/entity"
	subscription_model "awesomeProject/internal/domain/subscription/model"
	"context"
	"fmt"
)

type SubscriptionUseCase struct {
	subscriptionRepo subscription.Repository
}

func NewSubscriptionUseCase(subscriptionRepo subscription.Repository) *SubscriptionUseCase {
	return &SubscriptionUseCase{subscriptionRepo: subscriptionRepo}
}

func (uc *SubscriptionUseCase) CreateSubscription(ctx context.Context, dto subscription_model.CreateSubscriptionDTO) error {
	subscription := subscription_entity.NewSubscriptionFromCreate(dto.Title, dto.Duration, dto.Ruble, dto.Penny)
	if err := uc.subscriptionRepo.CreateSubscription(ctx, subscription); err != nil {
		return fmt.Errorf("SubscriptionUseCase - CreateSubscription - "+
			"subscriptionRepo.CreateSubscription: %w", err)
	}
	return nil
}
func (uc *SubscriptionUseCase) GetSubscription(ctx context.Context, subscriptionID int) (subscription_model.GetSubscriptionDTO, error) {
	subscription, err := uc.subscriptionRepo.GetSubscription(ctx, subscriptionID)
	if err != nil {
		return subscription_model.GetSubscriptionDTO{}, fmt.Errorf("SubscriptionUseCase - GetSubscription - "+
			"subscriptionRepo.GetSubscription: %w", err)
	}
	return subscription_model.NewGetSubscriptionResponse(subscription), nil
}
func (uc *SubscriptionUseCase) GetSubscriptionList(ctx context.Context) ([]subscription_model.GetSubscriptionDTO, error) {
	subscriptionList, err := uc.subscriptionRepo.GetSubscriptionList(ctx)
	if err != nil {
		return []subscription_model.GetSubscriptionDTO{}, fmt.Errorf("SubscriptionUseCase - "+
			"GetSubscriptionList - subscriptionRepo.GetSubscriptionList: %w", err)
	}
	return subscription_model.NewGetSubscriptionListResponse(subscriptionList), nil
}
func (uc *SubscriptionUseCase) UpdateSubscription(ctx context.Context, dto subscription_model.UpdateSubscriptionDTO) error {
	subscription := subscription_entity.NewSubscriptionFromUpdate(dto.Title, dto.ID, dto.Duration, dto.Ruble, dto.Penny)
	if err := uc.subscriptionRepo.UpdateSubscription(ctx, subscription); err != nil {
		return fmt.Errorf("SubscriptionUseCase - UpdateSubscription - subscriptionRepo.UpdateSubscription: %w", err)
	}
	return nil
}
func (uc *SubscriptionUseCase) DeleteSubscription(ctx context.Context, subscriptionID int) error {
	if err := uc.subscriptionRepo.DeleteSubscription(ctx, subscriptionID); err != nil {
		return fmt.Errorf("SubscriptionUseCase - DeleteSubscription - subscriptionRepo.DeleteSubscription: %w", err)
	}
	return nil
}
