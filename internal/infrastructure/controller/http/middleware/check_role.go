package middleware

import (
	auth_model "awesomeProject/internal/domain/auth/model"
	"awesomeProject/internal/infrastructure/controller/http/response"
	"github.com/gin-gonic/gin"
	"slices"
)

func CheckRole(roles []int) gin.HandlerFunc {
	return func(ctx *gin.Context) {

		member := ctx.MustGet("Member").(auth_model.MemberPayload)

		if !slices.Contains(roles, member.RoleID) {
			response.SendForbidden(ctx, "Нет доступа")
			return
		}
		// Продолжить выполнение
		ctx.Next()
	}
}
