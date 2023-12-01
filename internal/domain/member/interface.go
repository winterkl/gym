package member

import (
	"awesomeProject/internal/domain/member/entity"
	"awesomeProject/internal/domain/member/model"
	"context"
)

type UseCase interface {
	CreateMember(ctx context.Context, dto member_model.CreateMemberDTO) error
	GetMemberByLogin(ctx context.Context, login string) (member_model.GetMemberDTO, error)
	GetMember(ctx context.Context, memberID int) (member_model.GetMemberDTO, error)
	GetMemberList(ctx context.Context) ([]member_model.GetMemberDTO, error)
	UpdateMember(ctx context.Context, member member_model.UpdateMemberDTO) error
	DeleteMember(ctx context.Context, memberID int) error
	GetMemberByAuthData(ctx context.Context, login string, password string) (member_model.GetMemberDTO, error)
}

type Repository interface {
	CreateMember(ctx context.Context, member member_entity.Member) error
	GetMemberByLogin(ctx context.Context, login string) (member_entity.Member, error)
	GetMember(ctx context.Context, memberID int) (member_entity.Member, error)
	GetMemberList(ctx context.Context) ([]member_entity.Member, error)
	UpdateMember(ctx context.Context, member member_entity.Member) error
	DeleteMember(ctx context.Context, memberID int) error
	GetMemberByAuthData(ctx context.Context, login string, password string) (member_entity.Member, error)
}
