package auth_usecase

import (
	"awesomeProject/internal/app_errors"
	"awesomeProject/internal/domain/auth"
	"awesomeProject/internal/domain/auth/model"
	"awesomeProject/internal/domain/member"
	"awesomeProject/internal/domain/member/model"
	"context"
	"errors"
	"fmt"
)

type AuthUseCase struct {
	memberUC member.UseCase
	jwtAuth  auth.JwtAuth
}

func NewAuthUseCase(memberUC member.UseCase, jwtAuth auth.JwtAuth) *AuthUseCase {
	return &AuthUseCase{
		memberUC: memberUC,
		jwtAuth:  jwtAuth,
	}
}

func (uc *AuthUseCase) SignUp(ctx context.Context, dto auth_model.SignUpInputDTO) error {
	if _, err := uc.memberUC.GetMemberByLogin(ctx, dto.Login); err != nil {
		var errMemberNotFound *app_errors.MemberNotFound
		if !errors.As(err, &errMemberNotFound) {
			return fmt.Errorf("AuthUseCase - SignUp - memberUC.GetMemberByLogin: %w", err)
		}
	} else {
		return &app_errors.MemberAlreadyExists{Login: dto.Login}
	}
	createMemberDTO := member_model.NewCreateMemberDTO(dto.Login, dto.Password, dto.FIO)
	if err := uc.memberUC.CreateMember(ctx, createMemberDTO); err != nil {
		return err
	}
	return nil
}

func (uc *AuthUseCase) SignIn(ctx context.Context, dto auth_model.SignInInputDTO) (auth_model.SignInOutputDTO, error) {

	member, err := uc.memberUC.GetMemberByAuthData(ctx, dto.Login, dto.Password)
	if err != nil {
		var errMemberNotFound *app_errors.MemberNotFound
		if errors.As(err, &errMemberNotFound) {
			return auth_model.SignInOutputDTO{},
				fmt.Errorf("AuthUseCase - SignIn - memberUC.GetToken: %w", &app_errors.InvalidAuthData{})
		}
		return auth_model.SignInOutputDTO{}, fmt.Errorf("AuthUseCase - SignIn - memberUC.GetToken: %w", err)
	}

	token, err := uc.jwtAuth.GenerateToken(member.ID)
	if err != nil {
		return auth_model.SignInOutputDTO{}, fmt.Errorf("AuthUseCase - SignIn - jwtAuth.GenerateToken: %w", err)
	}

	return auth_model.NewSignInOutputDTO(member.ID, member.Login, member.FIO, token), nil
}
