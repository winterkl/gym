package v1

import (
	"awesomeProject/internal/app_errors"
	auth_model "awesomeProject/internal/domain/auth/model"
	member_entity "awesomeProject/internal/domain/member/entity"
	"awesomeProject/internal/domain/trainer"
	trainer_model "awesomeProject/internal/domain/trainer/model"
	"awesomeProject/internal/infrastructure/controller/http/middleware"
	"awesomeProject/internal/infrastructure/controller/http/response"
	"errors"
	"github.com/gin-gonic/gin"
	"strconv"
)

type trainerRoutes struct {
	trainerUC trainer.UseCase
}

func NewTrainerRouter(handler *gin.RouterGroup, trainerUC trainer.UseCase) {
	r := trainerRoutes{
		trainerUC: trainerUC,
	}
	trainerList := handler.Group("/trainer-list")
	{
		trainerList.POST("", r.createTrainer, middleware.CheckRole([]int{member_entity.RoleAdmin}))
		trainerList.GET("/:id", r.getTrainer)
		trainerList.GET("", r.getTrainerList)
		trainerList.DELETE("/:id", r.deleteTrainer, middleware.CheckRole([]int{member_entity.RoleAdmin}))
	}
}

func (r *trainerRoutes) createTrainer(ctx *gin.Context) {
	member := ctx.Value("Member").(auth_model.MemberPayload)
	trainerDTO := trainer_model.CreateTrainerDTO{
		MemberID: member.ID,
	}
	if err := r.trainerUC.CreateTrainer(ctx, trainerDTO); err != nil {
		handleTrainerError(ctx, err)
		return
	}
	response.SendOkRequest(ctx)
}
func (r *trainerRoutes) getTrainer(ctx *gin.Context) {
	trainerID, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		response.SendValidErrorRequest(ctx, err)
		return
	}
	trainer, err := r.trainerUC.GetTrainer(ctx, trainerID)
	if err != nil {
		handleTrainerError(ctx, err)
		return
	}
	response.SendOkRequestWithData(ctx, trainer)
}
func (r *trainerRoutes) getTrainerList(ctx *gin.Context) {
	trainerList, err := r.trainerUC.GetTrainerList(ctx)
	if err != nil {
		handleTrainerError(ctx, err)
		return
	}
	response.SendOkRequestWithData(ctx, trainerList)
}
func (r *trainerRoutes) deleteTrainer(ctx *gin.Context) {
	trainerID, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		response.SendValidErrorRequest(ctx, err)
		return
	}
	if err = r.trainerUC.DeleteTrainer(ctx, trainerID); err != nil {
		handleTrainerError(ctx, err)
		return
	}
	response.SendOkRequest(ctx)
}

func handleTrainerError(ctx *gin.Context, err error) {
	var errMemberAlreadyExists *app_errors.MemberAlreadyExists
	if errors.As(err, &errMemberAlreadyExists) {
		response.SendBadRequest(ctx, errMemberAlreadyExists.Error())
		return
	}
	var errMemberNotFound *app_errors.MemberNotFound
	if errors.As(err, &errMemberNotFound) {
		response.SendNotFound(ctx, errMemberNotFound.Error())
		return
	}
	response.SendInternalServerError(ctx, err)
}
