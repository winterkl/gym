package member

import (
	"awesomeProject/internal/domain/member/entity"
	"awesomeProject/internal/domain/member/model"
	"context"
)

type UseCase interface {
	CreateMember(ctx context.Context, dto model.CreateMemberDTO) error
	GetMemberByLogin(ctx context.Context, login string) (model.GetMemberDTO, error)
	GetMember(ctx context.Context, memberID int) (model.GetMemberDTO, error)
	GetMemberList(ctx context.Context) ([]model.GetMemberDTO, error)
	UpdateMember(ctx context.Context, member model.UpdateMemberDTO) error
	DeleteMember(ctx context.Context, memberID int) error
	GetMemberByAuthData(ctx context.Context, login string, password string) (model.GetMemberDTO, error)
}

type Repository interface {
	CreateMember(ctx context.Context, member entity.Member) error
	GetMemberByLogin(ctx context.Context, login string) (entity.Member, error)
	GetMember(ctx context.Context, memberID int) (entity.Member, error)
	GetMemberList(ctx context.Context) ([]entity.Member, error)
	UpdateMember(ctx context.Context, member entity.Member) error
	DeleteMember(ctx context.Context, memberID int) error
	GetMemberByAuthData(ctx context.Context, login string, password string) (entity.Member, error)
}
