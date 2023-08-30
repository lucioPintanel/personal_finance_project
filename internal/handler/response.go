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

type CreateTipoItemResponse struct {
	baseResponse
	Data int `json:"data"`
}

type DeleteTipoItemResponse struct {
	baseResponse
	Data int `json:"data"`
}

type ShowTipoItemResponse struct {
	baseResponse
	Data int `json:"data"`
}

type UpdateTipoItemResponse struct {
	baseResponse
	Data int `json:"data"`
}

type ListTipoItemsResponse struct {
	baseResponse
	Data []schemas.TipoItemResponse `json:"data"`
}

type CreateCardResponse struct {
	baseResponse
	Data int `json:"data"`
}

type DeleteCardResponse struct {
	baseResponse
	Data int `json:"data"`
}

type ShowCardResponse struct {
	baseResponse
	Data int `json:"data"`
}

type UpdateCardResponse struct {
	baseResponse
	Data int `json:"data"`
}

type ListCardsResponse struct {
	baseResponse
	Data []schemas.CardResponse `json:"data"`
}

type CreateUserResponse struct {
	baseResponse
	Data int `json:"data"`
}

type DeleteUserResponse struct {
	baseResponse
	Data int `json:"data"`
}

type ShowUserResponse struct {
	baseResponse
	Data int `json:"data"`
}

type UpdateUserResponse struct {
	baseResponse
	Data int `json:"data"`
}

type ListUsersResponse struct {
	baseResponse
	Data []schemas.UserResponse `json:"data"`
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
