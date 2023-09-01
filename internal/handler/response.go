package handler

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

type BaseResponse struct {
	Mensagem string `json:"mensagem"`
}

type ErrorResponse struct {
	BaseResponse
	ErroCode int `json:"errorCode"`
}

func SendError(ctx *gin.Context, statusCode int, msg string) {
	ctx.Header("Content-Type", "application/json")
	ctx.JSON(statusCode, gin.H{
		"mensagem":  msg,
		"errorCode": statusCode,
	})
}

func SendSuccess(ctx *gin.Context, op string, statusCode int, data interface{}) {
	ctx.Header("Content-Type", "application/json")
	ctx.JSON(statusCode, gin.H{
		"mensagem": fmt.Sprintf("operation from handler: %s successfull", op),
		"data":     data,
	})
}
