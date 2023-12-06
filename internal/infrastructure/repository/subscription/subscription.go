package subscription_repository

import (
	"awesomeProject/internal/app_errors"
	subscription_entity "awesomeProject/internal/domain/subscription/entity"
	"awesomeProject/pkg/postgres"
	"context"
	"database/sql"
	"errors"
	"fmt"
)

type SubscriptionRepository struct {
	db *postgres.Postgres
}

func NewSubscriptionRepository(db *postgres.Postgres) *SubscriptionRepository {
	return &SubscriptionRepository{
		db: db,
	}
}

func (r *SubscriptionRepository) CreateSubscription(ctx context.Context, subscription subscription_entity.Subscription) error {
	if err := r.db.
		NewInsert().
		Model(&subscription).
		Scan(ctx); err != nil {
		return fmt.Errorf("SubscriptionRepository - CreateSubscription - NewInsert: %w", err)
	}
	return nil
}
func (r *SubscriptionRepository) GetSubscription(ctx context.Context, subscriptionID int) (subscription_entity.Subscription, error) {
	subscription := subscription_entity.Subscription{}
	if err := r.db.
		NewSelect().
		Model(&subscription).
		Where("id = ?", subscriptionID).
		Scan(ctx); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return subscription, &app_errors.SubscriptionNotFound{ID: subscriptionID}
		}
		return subscription, fmt.Errorf("SubscriptionRepository - GetSubscription - NewSelect: %w", err)
	}
	return subscription, nil
}
func (r *SubscriptionRepository) GetSubscriptionList(ctx context.Context) ([]subscription_entity.Subscription, error) {
	subscriptionList := []subscription_entity.Subscription{}
	if err := r.db.
		NewSelect().
		Model(&subscriptionList).
		Scan(ctx); err != nil {
		return []subscription_entity.Subscription{}, fmt.Errorf("SubscriptionRepository - GetSubscriptionList - "+
			"NewSelect: %w", err)
	}
	return subscriptionList, nil
}
func (r *SubscriptionRepository) UpdateSubscription(ctx context.Context, subscription subscription_entity.Subscription) error {
	if _, err := r.db.
		NewUpdate().
		Model(&subscription).
		Set("title = ?", subscription.Title).
		Set("duration = ?", subscription.Duration).
		Set("ruble = ?", subscription.Ruble).
		Set("penny = ?", subscription.Penny).
		Where("id = ?", subscription.ID).
		Exec(ctx); err != nil {
		return fmt.Errorf("SubscriptionRepository - UpdateSubscription - NewUpdate: %w", err)
	}
	return nil
}
func (r *SubscriptionRepository) DeleteSubscription(ctx context.Context, subscriptionID int) error {
	if _, err := r.db.
		NewDelete().
		Model((*subscription_entity.Subscription)(nil)).
		Where("id = ?", subscriptionID).
		Exec(ctx); err != nil {
		return fmt.Errorf("SubscriptionRepository - DeleteSubscription - NewDelete: %w", err)
	}
	return nil
}
