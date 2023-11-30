package app_errors

type InvalidAuthData struct{}

func (a *InvalidAuthData) Error() string {
	return "Неверный логин или пароль"
}

type InvalidAccessToken struct{}

func (a *InvalidAccessToken) Error() string {
	return "Невалидный токен доступа"
}
