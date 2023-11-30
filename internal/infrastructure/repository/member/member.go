package member_repository

import (
	"awesomeProject/internal/app_errors"
	"awesomeProject/internal/domain/member/entity"
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

func (r *MemberRepository) CreateMember(ctx context.Context, member entity.Member) error {
	if err := r.db.NewInsert().Model(&member).Scan(ctx); err != nil {
		return fmt.Errorf("MemberRepository - CreateMember - NewInsert: %w", err)
	}
	return nil
}
func (r *MemberRepository) GetMemberByLogin(ctx context.Context, login string) (entity.Member, error) {
	var member entity.Member
	if err := r.db.NewSelect().Model(&member).Where("login = ?", login).Scan(ctx); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return member, &app_errors.MemberNotFound{Login: login}
		}
		return member, fmt.Errorf("MemberRepository - GetMemberByLogin - NewSelect: %w", err)
	}
	return member, nil
}
func (r *MemberRepository) GetMember(ctx context.Context, memberID int) (entity.Member, error) {
	member := entity.Member{}
	if err := r.db.NewSelect().Model(&member).Where("member.id = ?", memberID).Scan(ctx); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return member, &app_errors.MemberNotFound{ID: memberID}
		}
		return member, fmt.Errorf("MemberRepository - GetMember - NewSelect: %w", err)
	}
	return member, nil
}
func (r *MemberRepository) GetMemberByAuthData(ctx context.Context, login string, password string) (entity.Member, error) {
	var member entity.Member
	if err := r.db.NewSelect().Model(&member).
		Where("login = ?", login).
		Where("password = ?", password).Scan(ctx); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return member, &app_errors.MemberNotFound{Login: login}
		}
		return member, fmt.Errorf("MemberRepository - GetMember - NewSelect: %w", err)
	}
	return member, nil
}
func (r *MemberRepository) GetMemberList(ctx context.Context) ([]entity.Member, error) {
	memberList := []entity.Member{}
	if err := r.db.NewSelect().Model(&memberList).Scan(ctx); err != nil {
		return []entity.Member{}, fmt.Errorf("MemberRepository - GetMemberList - NewSelect: %w", err)
	}
	return memberList, nil
}
func (r *MemberRepository) UpdateMember(ctx context.Context, member entity.Member) error {
	if _, err := r.db.NewUpdate().Model(&member).
		Set("fio = ?", member.FIO).Set("password = ?", member.Password).
		Where("id = ?", member.ID).Exec(ctx); err != nil {
		return fmt.Errorf("MemberRepository - UpdateMember - NewUpdate: %w", err)
	}
	return nil
}
func (r *MemberRepository) DeleteMember(ctx context.Context, memberID int) error {
	if _, err := r.db.NewDelete().Model((*entity.Member)(nil)).
		Where("id = ?", memberID).Exec(ctx); err != nil {
		return fmt.Errorf("MemberRepository - DeletMember - NewDelete: %w", err)
	}
	return nil
}
