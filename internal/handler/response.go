package handler

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lucioPintanel/personal_finance_project/internal/schemas"
)

type baseResponse struct {
	Mensagem string `json:"mensagem"`
}

type ErrorResponse struct {
	baseResponse
	ErroCode int `json:"errorCode"`
}

type CreateTipoResponse struct {
	baseResponse
	Data int `json:"data"`
}

type DeleteTipoResponse struct {
	baseResponse
	Data int `json:"data"`
}

type ShowTipoResponse struct {
	baseResponse
	Data int `json:"data"`
}

type UpdateTipoResponse struct {
	baseResponse
	Data int `json:"data"`
}

type ListTipoResponse struct {
	baseResponse
	Data []schemas.TipoResponse `json:"data"`
}

func SendError(ctx *gin.Context, code int, msg string) {
	ctx.Header("Content-Type", "application/json")
	ctx.JSON(code, gin.H{
		"mensagem":  msg,
		"errorCode": code,
	})
}

func SendSuccess(ctx *gin.Context, op string, data interface{}) {
	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, gin.H{
		"mensagem": fmt.Sprintf("operation from handler: %s successfull", op),
		"data":     data,
	})
}
