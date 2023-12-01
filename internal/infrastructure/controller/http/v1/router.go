package v1

import (
	"awesomeProject/internal/domain/auth"
	"awesomeProject/internal/domain/member"
	"awesomeProject/internal/domain/trainer"
	v1 "awesomeProject/internal/infrastructure/controller/http/middleware"
	"github.com/gin-gonic/gin"
)

type UC struct {
	MemberUC  member.UseCase
	TrainerUC trainer.UseCase
	AuthUC    auth.UseCase
}

func NewRouter(handler *gin.Engine, uc UC, jwtAuth auth.JwtAuth) {
	h := handler.Group("v1")

	NewAuthRoutes(h, uc.AuthUC)
	h.Use(v1.ParseToken(jwtAuth, uc.MemberUC))
	{
		NewTrainerRouter(h, uc.TrainerUC)
		NewMemberRouter(h, uc.MemberUC)
	}
}
