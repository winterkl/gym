package service_entity

type Service struct {
	ID       int `bun:"id,pk,autoincrement"`
	Title    string
	Duration int
	Ruble    int
	Penny    int
}

func NewServiceFromCreate(title string, duration, ruble, penny int) Service {
	return Service{
		Title:    title,
		Duration: duration,
		Ruble:    ruble,
		Penny:    penny,
	}
}
func NewServiceFromUpdate(title string, serviceID, duration, ruble, penny int) Service {
	return Service{
		ID:       serviceID,
		Title:    title,
		Duration: duration,
		Ruble:    ruble,
		Penny:    penny,
	}
}
