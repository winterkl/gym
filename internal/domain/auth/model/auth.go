package auth_model

type MemberPayload struct {
	ID     int    `json:"id"`
	Login  string `json:"login"`
	FIO    string `json:"fio"`
	RoleID int    `json:"role_id"`
}

func NewMemberPayload(id, roleID int, login, fio string) MemberPayload {
	return MemberPayload{
		ID:     id,
		Login:  login,
		FIO:    fio,
		RoleID: roleID,
	}
}

type SignInInputDTO struct {
	Login    string `json:"login" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type SignInOutputDTO struct {
	Token         string        `json:"token"`
	MemberPayload MemberPayload `json:"member_payload"`
}

func NewSignInOutputDTO(memberID, roleID int, login, fio, token string) SignInOutputDTO {
	return SignInOutputDTO{
		Token:         token,
		MemberPayload: NewMemberPayload(memberID, roleID, login, fio),
	}
}

type SignUpInputDTO struct {
	Login    string `json:"login" binding:"required"`
	Password string `json:"password" binding:"required"`
	FIO      string `json:"fio" binding:"required"`
}
