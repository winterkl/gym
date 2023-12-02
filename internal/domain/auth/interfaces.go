package auth

import (
	auth_model "awesomeProject/internal/domain/auth/model"
	"context"
)

type UseCase interface {
	SignUp(ctx context.Context, dto auth_model.SignUpInputDTO) error
	SignIn(ctx context.Context, dto auth_model.SignInInputDTO) (auth_model.SignInOutputDTO, error)
}
