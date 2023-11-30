package response

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"os"
)

const (
	MIME_AUDIO = "audio/mpeg"
)

type Response struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

func failureResponse(ctx *gin.Context, status int, response Response) {
	ctx.AbortWithStatusJSON(status, response)
}

func successResponse(ctx *gin.Context, status int, response Response) {
	ctx.JSON(status, response)
}

func SendInternalServerError(ctx *gin.Context, err error) {
	log.Println(err)
	failureResponse(ctx, http.StatusInternalServerError, Response{
		Message: "Internal Server Error",
		Data:    nil,
	})
}

func SendBadRequest(ctx *gin.Context, message string) {
	failureResponse(ctx, http.StatusBadRequest, Response{
		Message: message,
		Data:    nil,
	})
}

func SendNotFound(ctx *gin.Context, message string) {
	failureResponse(ctx, http.StatusNotFound, Response{
		Message: message,
		Data:    nil,
	})
}

func SendUnauthorized(ctx *gin.Context, message string) {
	failureResponse(ctx, http.StatusUnauthorized, Response{
		Message: message,
		Data:    nil,
	})
}

func SendForbidden(ctx *gin.Context, message string) {
	failureResponse(ctx, http.StatusForbidden, Response{
		Message: message,
		Data:    nil,
	})
}

func SendOkRequest(ctx *gin.Context) {
	successResponse(ctx, http.StatusOK, Response{
		Message: "ok",
		Data:    nil,
	})
}

func SendOkRequestWithData(ctx *gin.Context, data interface{}) {
	successResponse(ctx, http.StatusOK, Response{
		Message: "ok",
		Data:    &data,
	})
}

func SendOkRequestWithFile(ctx *gin.Context, file *os.File, fileName, contentType string) {
	ctx.Header("Content-Disposition", "attachment; filename="+fileName)
	ctx.Header("Content-Description", "File Transfer")
	ctx.Header("Content-Type", contentType)
	ctx.File(file.Name())
}

func SendValidErrorRequest(ctx *gin.Context, data interface{}) {
	failureResponse(ctx, http.StatusBadRequest, Response{
		Message: "validation error",
		Data:    data,
	})
}
