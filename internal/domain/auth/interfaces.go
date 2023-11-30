package auth

import (
	"awesomeProject/internal/domain/auth/model"
	"context"
)

type UseCase interface {
	SignUp(ctx context.Context, dto model.SignUpInputDTO) error
	SignIn(ctx context.Context, dto model.SignInInputDTO) (model.SignInOutputDTO, error)
}
