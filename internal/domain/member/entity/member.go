package entity

type Member struct {
	ID       int `bun:"id,pk,autoincrement"`
	Login    string
	Password string
	FIO      string
}

func NewMemberFromCreate(login, password, fio string) Member {
	return Member{
		Login:    login,
		Password: password,
		FIO:      fio,
	}
}
func NewMemberFromUpdate(memberID int, fio, password string) Member {
	return Member{
		ID:       memberID,
		Password: password,
		FIO:      fio,
	}
}
