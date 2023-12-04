package member_usecase

import (
	"awesomeProject/internal/domain/member"
	member_entity "awesomeProject/internal/domain/member/entity"
	member_model "awesomeProject/internal/domain/member/model"
	"context"
	"fmt"
)

type MemberUseCase struct {
	memberRepo member.Repository
}

func NewMemberUseCase(memberRepo member.Repository) *MemberUseCase {
	return &MemberUseCase{memberRepo: memberRepo}
}

func (uc *MemberUseCase) CreateMember(ctx context.Context, dto member_model.CreateMemberDTO) error {
	member := member_entity.NewMemberFromCreate(dto.Login, dto.Password, dto.FIO)
	if err := uc.memberRepo.CreateMember(ctx, member); err != nil {
		return fmt.Errorf("MemberUseCase - CreateMember - memberRepo.CreateMember: %w", err)
	}
	return nil
}
func (uc *MemberUseCase) GetMemberByLogin(ctx context.Context, login string) (member_model.GetMemberDTO, error) {
	member, err := uc.memberRepo.GetMemberByLogin(ctx, login)
	if err != nil {
		return member_model.GetMemberDTO{}, fmt.Errorf("MemberUseCase - GetMemberByLogin - "+
			"memberRepo.GetMemberByLogin: %w", err)
	}
	return member_model.NewGetMemberResponse(member), nil
}
func (uc *MemberUseCase) GetMember(ctx context.Context, memberID int) (member_model.GetMemberDTO, error) {
	member, err := uc.memberRepo.GetMember(ctx, memberID)
	if err != nil {
		return member_model.GetMemberDTO{}, fmt.Errorf("MemberUseCase - GetMember - memberRepo.GetMember: %w", err)
	}
	return member_model.NewGetMemberResponse(member), nil
}
func (uc *MemberUseCase) GetMemberList(ctx context.Context) ([]member_model.GetMemberDTO, error) {
	memberList, err := uc.memberRepo.GetMemberList(ctx)
	if err != nil {
		return []member_model.GetMemberDTO{}, fmt.Errorf("MemberUseCase - GetMemberList - "+
			"memberRepo.GetMemberList: %w", err)
	}
	return member_model.NewGetMemberListResponse(memberList), nil
}
func (uc *MemberUseCase) UpdateMember(ctx context.Context, dto member_model.UpdateMemberDTO) error {
	member := member_entity.NewMemberFromUpdate(dto.ID, dto.Password, dto.FIO)
	if err := uc.memberRepo.UpdateMember(ctx, member); err != nil {
		return fmt.Errorf("MemberUseCase - UpdateMember - memberRepo.UpdateMember: %w", err)
	}
	return nil
}
func (uc *MemberUseCase) DeleteMember(ctx context.Context, memberID int) error {
	if err := uc.memberRepo.DeleteMember(ctx, memberID); err != nil {
		return fmt.Errorf("MemberUseCase - DeleteMember - memberRepo.DeleteMember: %w", err)
	}
	return nil
}
func (uc *MemberUseCase) GetMemberByAuthData(ctx context.Context, login string, password string) (member_model.GetMemberDTO, error) {
	member, err := uc.memberRepo.GetMemberByAuthData(ctx, login, password)
	if err != nil {
		return member_model.GetMemberDTO{}, fmt.Errorf("MemberUseCase - GetMemberByLogin - "+
			"memberRepo.GetMemberByAuthData: %w", err)
	}
	return member_model.NewGetMemberResponse(member), nil
}

// Метод для админа

func (uc *MemberUseCase) UpdateRole(ctx context.Context, dto member_model.UpdateRoleDTO) error {
	if err := uc.memberRepo.UpdateRole(ctx, dto.ID, dto.RoleID); err != nil {
		return fmt.Errorf("MemberUseCase - UpdateRole - memberRepo.UpdateRole: %w", err)
	}
	return nil
}
