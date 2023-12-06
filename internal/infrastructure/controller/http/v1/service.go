package v1

import (
	"awesomeProject/internal/app_errors"
	member_entity "awesomeProject/internal/domain/member/entity"
	"awesomeProject/internal/domain/service"
	service_model "awesomeProject/internal/domain/service/model"
	"awesomeProject/internal/infrastructure/controller/http/middleware"
	"awesomeProject/internal/infrastructure/controller/http/response"
	"errors"
	"github.com/gin-gonic/gin"
	"strconv"
)

type serviceRoutes struct {
	serviceUC service.UseCase
}

func NewServiceRouter(handler *gin.RouterGroup, serviceUC service.UseCase) {
	r := serviceRoutes{
		serviceUC: serviceUC,
	}
	serviceList := handler.Group("/service-list", middleware.CheckRole([]int{member_entity.RoleAdmin}))
	{
		serviceList.POST("", r.createService)
		serviceList.GET("", r.getServiceList)
		serviceList.GET("/:id", r.getService)
		serviceList.PUT("/:id", r.updateService)
		serviceList.DELETE("/:id", r.deleteService)
	}
}

func (r *serviceRoutes) createService(ctx *gin.Context) {
	serviceDTO := service_model.CreateServiceDTO{}
	if err := ctx.ShouldBindJSON(&serviceDTO); err != nil {
		response.SendValidErrorRequest(ctx, err.Error())
		return
	}
	if err := r.serviceUC.CreateService(ctx, serviceDTO); err != nil {
		handleServiceError(ctx, err)
		return
	}
	response.SendOkRequest(ctx)
}
func (r *serviceRoutes) getService(ctx *gin.Context) {
	serviceID, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		response.SendValidErrorRequest(ctx, err)
		return
	}
	service, err := r.serviceUC.GetService(ctx, serviceID)
	if err != nil {
		handleServiceError(ctx, err)
		return
	}
	response.SendOkRequestWithData(ctx, service)
}
func (r *serviceRoutes) getServiceList(ctx *gin.Context) {
	serviceList, err := r.serviceUC.GetServiceList(ctx)
	if err != nil {
		handleServiceError(ctx, err)
		return
	}
	response.SendOkRequestWithData(ctx, serviceList)
}
func (r *serviceRoutes) updateService(ctx *gin.Context) {
	serviceID, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		response.SendValidErrorRequest(ctx, err)
		return
	}
	serviceDTO := service_model.UpdateServiceDTO{ID: serviceID}
	if err = ctx.ShouldBindJSON(&serviceDTO); err != nil {
		response.SendValidErrorRequest(ctx, err.Error())
		return
	}
	if err = r.serviceUC.UpdateService(ctx, serviceDTO); err != nil {
		handleServiceError(ctx, err)
		return
	}
	response.SendOkRequest(ctx)
}
func (r *serviceRoutes) deleteService(ctx *gin.Context) {
	serviceID, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		response.SendValidErrorRequest(ctx, err)
		return
	}
	if err = r.serviceUC.DeleteService(ctx, serviceID); err != nil {
		handleServiceError(ctx, err)
		return
	}
	response.SendOkRequest(ctx)
}

func handleServiceError(ctx *gin.Context, err error) {
	var errServiceAlreadyExists *app_errors.ServiceAlreadyExists
	if errors.As(err, &errServiceAlreadyExists) {
		response.SendBadRequest(ctx, errServiceAlreadyExists.Error())
		return
	}
	var errServiceNotFound *app_errors.ServiceNotFound
	if errors.As(err, &errServiceNotFound) {
		response.SendNotFound(ctx, errServiceNotFound.Error())
		return
	}
	response.SendInternalServerError(ctx, err)
}
