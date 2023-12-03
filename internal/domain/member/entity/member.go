package member_entity

type Member struct {
	ID       int `bun:"id,pk,autoincrement"`
	Login    string
	Password string
	FIO      string
	RoleID   int
	Role     Role `bun:"rel:belongs-to,join:role_id=id"`
}

func NewMemberFromCreate(login, password, fio string) Member {
	return Member{
		Login:    login,
		Password: password,
		FIO:      fio,
		RoleID:   RoleMember,
	}
}
func NewMemberFromUpdate(memberID int, password, fio string) Member {
	return Member{
		ID:       memberID,
		Password: password,
		FIO:      fio,
	}
}
