package v1

import (
	"awesomeProject/internal/app_errors"
	"awesomeProject/internal/domain/auth"
	"awesomeProject/internal/domain/auth/model"
	"awesomeProject/internal/infrastructure/controller/http/response"
	"errors"
	"github.com/gin-gonic/gin"
)

type authRoutes struct {
	authUC auth.UseCase
}

func NewAuthRoutes(handler *gin.RouterGroup, authUC auth.UseCase) {
	r := authRoutes{
		authUC: authUC,
	}
	auth := handler.Group("/auth")
	{
		auth.POST("/sign-in", r.signIn)
		auth.POST("/sign-up", r.signUp)
	}
}

func (r *authRoutes) signIn(ctx *gin.Context) {
	userDTO := auth_model.SignInInputDTO{}
	if err := ctx.ShouldBindJSON(&userDTO); err != nil {
		response.SendValidErrorRequest(ctx, err.Error())
		return
	}

	authData, err := r.authUC.SignIn(ctx, userDTO)
	if err != nil {
		response.SendValidErrorRequest(ctx, err.Error())
		return
	}
	response.SendOkRequestWithData(ctx, authData)
}
func (r *authRoutes) signUp(ctx *gin.Context) {
	userDTO := auth_model.SignUpInputDTO{}
	if err := ctx.ShouldBindJSON(&userDTO); err != nil {
		response.SendValidErrorRequest(ctx, err.Error())
		return
	}
	if err := r.authUC.SignUp(ctx, userDTO); err != nil {
		handleAuthError(ctx, err)
		return
	}
	response.SendOkRequest(ctx)
}

func handleAuthError(ctx *gin.Context, err error) {
	var errUserAlreadyExists *app_errors.MemberAlreadyExists
	if errors.As(err, &errUserAlreadyExists) {
		response.SendNotFound(ctx, errUserAlreadyExists.Error())
		return
	}
	response.SendInternalServerError(ctx, err)
}
