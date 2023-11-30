package auth

type JwtAuth interface {
	GenerateToken(userID int) (string, error)
	ParseToken(tokenStr string) (int, error)
}
