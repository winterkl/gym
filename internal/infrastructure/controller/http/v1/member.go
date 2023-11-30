package v1

import (
	"awesomeProject/internal/app_errors"
	"awesomeProject/internal/domain/member"
	"awesomeProject/internal/domain/member/model"
	"awesomeProject/internal/infrastructure/controller/http/response"
	"errors"
	"github.com/gin-gonic/gin"
	"strconv"
)

type memberRoutes struct {
	memberUC member.UseCase
}

func NewMemberRouter(handler *gin.RouterGroup, memberUC member.UseCase) {
	r := memberRoutes{
		memberUC: memberUC,
	}
	memberList := handler.Group("/member-list")
	{
		memberList.POST("", r.createMember)
		memberList.GET("", r.getMemberList)
		memberList.GET("/:id", r.getMember)
		memberList.PUT("/:id", r.updateMember)
		memberList.DELETE("/:id", r.deleteMember)
	}
}

func (r *memberRoutes) createMember(ctx *gin.Context) {
	memberDTO := model.CreateMemberDTO{}
	if err := ctx.ShouldBindJSON(&memberDTO); err != nil {
		response.SendValidErrorRequest(ctx, err.Error())
		return
	}
	if err := r.memberUC.CreateMember(ctx, memberDTO); err != nil {
		handleMemberError(ctx, err)
		return
	}
	response.SendOkRequest(ctx)
}
func (r *memberRoutes) getMember(ctx *gin.Context) {
	memberID, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		response.SendValidErrorRequest(ctx, err)
		return
	}
	member, err := r.memberUC.GetMember(ctx, memberID)
	if err != nil {
		handleMemberError(ctx, err)
		return
	}
	response.SendOkRequestWithData(ctx, member)
}
func (r *memberRoutes) getMemberList(ctx *gin.Context) {
	memberList, err := r.memberUC.GetMemberList(ctx)
	if err != nil {
		handleMemberError(ctx, err)
		return
	}
	response.SendOkRequestWithData(ctx, memberList)
}
func (r *memberRoutes) updateMember(ctx *gin.Context) {
	memberID, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		response.SendValidErrorRequest(ctx, err)
		return
	}
	memberDTO := model.UpdateMemberDTO{ID: memberID}
	if err = ctx.ShouldBindJSON(&memberDTO); err != nil {
		response.SendValidErrorRequest(ctx, err.Error())
		return
	}
	if err = r.memberUC.UpdateMember(ctx, memberDTO); err != nil {
		handleMemberError(ctx, err)
		return
	}
	response.SendOkRequest(ctx)
}
func (r *memberRoutes) deleteMember(ctx *gin.Context) {
	memberID, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		response.SendValidErrorRequest(ctx, err)
		return
	}
	if err = r.memberUC.DeleteMember(ctx, memberID); err != nil {
		handleMemberError(ctx, err)
		return
	}
	response.SendOkRequest(ctx)
}

func handleMemberError(ctx *gin.Context, err error) {
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
