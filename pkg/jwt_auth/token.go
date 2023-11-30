package jwt_auth

import (
	"awesomeProject/internal/app_errors"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"time"
)

type JwtAuth struct {
	key string
}

func NewJwtAuth(key string) *JwtAuth {
	return &JwtAuth{
		key: key,
	}
}

type TokenClaims struct {
	jwt.StandardClaims
	MemberID int
}

func (ja *JwtAuth) GenerateToken(memberID int) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &TokenClaims{
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(12 * time.Hour).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
		memberID,
	})

	return token.SignedString([]byte(ja.key))
}

func (ja *JwtAuth) ParseToken(tokenStr string) (int, error) {
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		return []byte(ja.key), nil
	})
	if err != nil {
		return 0, fmt.Errorf("JwtAuth - ParseToken - jwt.Parse: %w", err)
	}
	if !token.Valid {
		return 0, &app_errors.InvalidAccessToken{}
	}
	return int(token.Claims.(jwt.MapClaims)["MemberID"].(float64)), nil
}
