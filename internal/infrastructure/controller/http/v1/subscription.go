package v1

import (
	"awesomeProject/internal/app_errors"
	member_entity "awesomeProject/internal/domain/member/entity"
	"awesomeProject/internal/domain/subscription"
	subscription_model "awesomeProject/internal/domain/subscription/model"
	"awesomeProject/internal/infrastructure/controller/http/middleware"
	"awesomeProject/internal/infrastructure/controller/http/response"
	"errors"
	"github.com/gin-gonic/gin"
	"strconv"
)

type subscriptionRoutes struct {
	subscriptionUC subscription.UseCase
}

func NewSubscriptionRouter(handler *gin.RouterGroup, subscriptionUC subscription.UseCase) {
	r := subscriptionRoutes{
		subscriptionUC: subscriptionUC,
	}
	subscriptionList := handler.Group("/subscription-list", middleware.CheckRole([]int{member_entity.RoleAdmin}))
	{
		subscriptionList.POST("", r.createSubscription)
		subscriptionList.GET("", r.getSubscriptionList)
		subscriptionList.GET("/:id", r.getSubscription)
		subscriptionList.PUT("/:id", r.updateSubscription)
		subscriptionList.DELETE("/:id", r.deleteSubscription)
	}
}

func (r *subscriptionRoutes) createSubscription(ctx *gin.Context) {
	subscriptionDTO := subscription_model.CreateSubscriptionDTO{}
	if err := ctx.ShouldBindJSON(&subscriptionDTO); err != nil {
		response.SendValidErrorRequest(ctx, err.Error())
		return
	}
	if err := r.subscriptionUC.CreateSubscription(ctx, subscriptionDTO); err != nil {
		handleSubscriptionError(ctx, err)
		return
	}
	response.SendOkRequest(ctx)
}
func (r *subscriptionRoutes) getSubscription(ctx *gin.Context) {
	subscriptionID, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		response.SendValidErrorRequest(ctx, err)
		return
	}
	subscription, err := r.subscriptionUC.GetSubscription(ctx, subscriptionID)
	if err != nil {
		handleSubscriptionError(ctx, err)
		return
	}
	response.SendOkRequestWithData(ctx, subscription)
}
func (r *subscriptionRoutes) getSubscriptionList(ctx *gin.Context) {
	subscriptionList, err := r.subscriptionUC.GetSubscriptionList(ctx)
	if err != nil {
		handleSubscriptionError(ctx, err)
		return
	}
	response.SendOkRequestWithData(ctx, subscriptionList)
}
func (r *subscriptionRoutes) updateSubscription(ctx *gin.Context) {
	subscriptionID, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		response.SendValidErrorRequest(ctx, err)
		return
	}
	subscriptionDTO := subscription_model.UpdateSubscriptionDTO{ID: subscriptionID}
	if err = ctx.ShouldBindJSON(&subscriptionDTO); err != nil {
		response.SendValidErrorRequest(ctx, err.Error())
		return
	}
	if err = r.subscriptionUC.UpdateSubscription(ctx, subscriptionDTO); err != nil {
		handleSubscriptionError(ctx, err)
		return
	}
	response.SendOkRequest(ctx)
}
func (r *subscriptionRoutes) deleteSubscription(ctx *gin.Context) {
	subscriptionID, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		response.SendValidErrorRequest(ctx, err)
		return
	}
	if err = r.subscriptionUC.DeleteSubscription(ctx, subscriptionID); err != nil {
		handleSubscriptionError(ctx, err)
		return
	}
	response.SendOkRequest(ctx)
}

func handleSubscriptionError(ctx *gin.Context, err error) {
	var errSubscriptionAlreadyExists *app_errors.SubscriptionAlreadyExists
	if errors.As(err, &errSubscriptionAlreadyExists) {
		response.SendBadRequest(ctx, errSubscriptionAlreadyExists.Error())
		return
	}
	var errSubscriptionNotFound *app_errors.SubscriptionNotFound
	if errors.As(err, &errSubscriptionNotFound) {
		response.SendNotFound(ctx, errSubscriptionNotFound.Error())
		return
	}
	response.SendInternalServerError(ctx, err)
}
