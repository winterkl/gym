package middleware

import (
	"awesomeProject/internal/app_errors"
	"awesomeProject/internal/domain/auth"
	"awesomeProject/internal/domain/auth/model"
	"awesomeProject/internal/domain/member"
	"awesomeProject/internal/infrastructure/controller/http/response"
	"errors"
	"github.com/gin-gonic/gin"
	"strings"
)

func ParseToken(jwtAuth auth.JwtAuth, memberUC member.UseCase) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token, err := getTokenFromHeader(ctx)
		if err != nil {
			response.SendUnauthorized(ctx, err.Error())
			return
		}
		memberID, err := jwtAuth.ParseToken(token)
		if err != nil {
			response.SendUnauthorized(ctx, err.Error())
			return
		}

		// Ищем участника в нашей базе
		member, err := memberUC.GetMember(ctx, memberID)
		if err != nil {
			var errMemberNotFound *app_errors.MemberNotFound
			if errors.As(err, &errMemberNotFound) {
				response.SendUnauthorized(ctx, errMemberNotFound.Error())
				return
			}
			response.SendInternalServerError(ctx, err)
			return
		}

		memberPayload := auth_model.NewMemberPayload(member.ID, member.Login, member.FIO)
		ctx.Set("Member", memberPayload)

		ctx.Next()
	}
}

func getTokenFromHeader(ctx *gin.Context) (string, error) {
	authHeader := ctx.Request.Header.Get("Authorization")
	if authHeader == "" {
		return "", errors.New("authorization header not set")
	}

	bearerTokenParts := strings.Split(authHeader, "Bearer")
	if len(bearerTokenParts) < 2 {
		return "", errors.New("authorization header has wrong format")
	}

	token := strings.TrimSpace(bearerTokenParts[1])
	if token == "" {
		return "", errors.New("authorization token not set")
	}

	return token, nil
}
