package member_repository

import (
	"awesomeProject/internal/app_errors"
	member_entity "awesomeProject/internal/domain/member/entity"
	"awesomeProject/pkg/postgres"
	"context"
	"database/sql"
	"errors"
	"fmt"
)

type MemberRepository struct {
	db *postgres.Postgres
}

func NewMemberRepository(db *postgres.Postgres) *MemberRepository {
	return &MemberRepository{
		db: db,
	}
}

func (r *MemberRepository) CreateMember(ctx context.Context, member member_entity.Member) error {
	if err := r.db.
		NewInsert().
		Model(&member).
		Scan(ctx); err != nil {
		return fmt.Errorf("MemberRepository - CreateMember - NewInsert: %w", err)
	}
	return nil
}
func (r *MemberRepository) GetMemberByLogin(ctx context.Context, login string) (member_entity.Member, error) {
	var member member_entity.Member
	if err := r.db.
		NewSelect().
		Model(&member).
		Where("login = ?", login).
		Scan(ctx); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return member, &app_errors.MemberNotFound{Login: login}
		}
		return member, fmt.Errorf("MemberRepository - GetMemberByLogin - NewSelect: %w", err)
	}
	return member, nil
}
func (r *MemberRepository) GetMember(ctx context.Context, memberID int) (member_entity.Member, error) {
	member := member_entity.Member{}
	if err := r.db.
		NewSelect().
		Model(&member).
		Where("member.id = ?", memberID).
		Scan(ctx); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return member, &app_errors.MemberNotFound{ID: memberID}
		}
		return member, fmt.Errorf("MemberRepository - GetMember - NewSelect: %w", err)
	}
	return member, nil
}
func (r *MemberRepository) GetMemberByAuthData(ctx context.Context, login string, password string) (member_entity.Member, error) {
	var member member_entity.Member
	if err := r.db.
		NewSelect().
		Model(&member).
		Where("login = ?", login).
		Where("password = ?", password).
		Scan(ctx); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return member, &app_errors.MemberNotFound{Login: login}
		}
		return member, fmt.Errorf("MemberRepository - GetMember - NewSelect: %w", err)
	}
	return member, nil
}
func (r *MemberRepository) GetMemberList(ctx context.Context) ([]member_entity.Member, error) {
	memberList := []member_entity.Member{}
	if err := r.db.
		NewSelect().
		Model(&memberList).
		Scan(ctx); err != nil {
		return []member_entity.Member{}, fmt.Errorf("MemberRepository - GetMemberList - NewSelect: %w", err)
	}
	return memberList, nil
}
func (r *MemberRepository) UpdateMember(ctx context.Context, member member_entity.Member) error {
	if _, err := r.db.
		NewUpdate().
		Model(&member).
		Set("fio = ?", member.FIO).
		Set("password = ?", member.Password).
		Where("id = ?", member.ID).
		Exec(ctx); err != nil {
		return fmt.Errorf("MemberRepository - UpdateMember - NewUpdate: %w", err)
	}
	return nil
}
func (r *MemberRepository) DeleteMember(ctx context.Context, memberID int) error {
	if _, err := r.db.
		NewDelete().
		Model((*member_entity.Member)(nil)).
		Where("id = ?", memberID).
		Exec(ctx); err != nil {
		return fmt.Errorf("MemberRepository - DeleteMember - NewDelete: %w", err)
	}
	return nil
}

// Метод для админа

func (r *MemberRepository) UpdateRole(ctx context.Context, ID, roleID int) error {
	if _, err := r.db.
		NewUpdate().
		Model((*member_entity.Member)(nil)).
		Set("role_id = ?", roleID).
		Where("id = ?", ID).
		Exec(ctx); err != nil {
		return fmt.Errorf("MemberRepository - UpdateRole - NewUpdate: %w", err)
	}
	return nil
}
