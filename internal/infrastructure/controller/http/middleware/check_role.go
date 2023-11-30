package middleware

//
//import (
//	"github.com/MelethFaer/shop-manager/internal/domain/auth/model"
//	"github.com/MelethFaer/shop-manager/internal/infrastructure/controller/http/response"
//	"github.com/gin-gonic/gin"
//	"slices"
//)
//
//func CheckRole(roles []int) gin.HandlerFunc {
//	return func(ctx *gin.Context) {
//
//		user := ctx.MustGet("User").(model.UserPayload)
//
//		if !slices.Contains(roles, user.RoleID) {
//			response.SendForbidden(ctx, "Нет доступа")
//			return
//		}
//		// Продолжить выполнение
//		ctx.Next()
//	}
//}
