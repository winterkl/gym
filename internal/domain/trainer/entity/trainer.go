package trainer_entity

import member_entity "awesomeProject/internal/domain/member/entity"

type Trainer struct {
	ID       int `bun:"id,pk,autoincrement"`
	MemberID int
	Member   member_entity.Member `bun:"rel:belongs-to,join:member_id=id"`
}

func NewTrainerFromCreate(memberID int) Trainer {
	return Trainer{
		MemberID: memberID,
	}
}
