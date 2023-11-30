package model

import "awesomeProject/internal/domain/member/entity"

type CreateMemberDTO struct {
	Login    string `json:"login" binding:"required"`
	Password string `json:"password" binding:"required"`
	FIO      string `json:"fio" binding:"required"`
}

type GetMemberDTO struct {
	ID    int    `json:"id"`
	Login string `json:"login"`
	FIO   string `json:"fio"`
}

func NewCreateMemberDTO(login, password, fio string) CreateMemberDTO {
	return CreateMemberDTO{
		Login:    login,
		Password: password,
		FIO:      fio,
	}
}
func NewGetMemberResponse(member entity.Member) GetMemberDTO {
	return GetMemberDTO{
		ID:    member.ID,
		Login: member.Login,
		FIO:   member.FIO,
	}
}
func NewGetMemberListResponse(memberList []entity.Member) []GetMemberDTO {
	memberListDTO := []GetMemberDTO{}
	for _, member := range memberList {
		memberListDTO = append(memberListDTO, NewGetMemberResponse(member))
	}
	return memberListDTO
}

type UpdateMemberDTO struct {
	ID       int
	Password string `json:"password"`
	FIO      string `json:"fio"`
}
