package subscription_entity

type Subscription struct {
	ID       int `bun:"id,pk,autoincrement"`
	Title    string
	Duration int
	Ruble    int
	Penny    int
}

func NewSubscriptionFromCreate(title string, duration, ruble, penny int) Subscription {
	return Subscription{
		Title:    title,
		Duration: duration,
		Ruble:    ruble,
		Penny:    penny,
	}
}
func NewSubscriptionFromUpdate(title string, subscriptionID, duration, ruble, penny int) Subscription {
	return Subscription{
		ID:       subscriptionID,
		Title:    title,
		Duration: duration,
		Ruble:    ruble,
		Penny:    penny,
	}
}
