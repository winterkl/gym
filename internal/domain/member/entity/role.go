package member_entity

const (
	RoleAdmin   = 1
	RoleTrainer = 2
	RoleMember  = 3
)

var roleList = map[int]string{
	RoleAdmin:   "Администратор",
	RoleTrainer: "Тренер",
	RoleMember:  "Участник",
}

type Role struct {
	ID       int `bun:"id,pk,autoincrement"`
	CodeName string
	Title    string
}
