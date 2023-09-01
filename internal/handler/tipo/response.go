package tipo

import (
	"github.com/lucioPintanel/personal_finance_project/internal/handler"
	"github.com/lucioPintanel/personal_finance_project/internal/schemas"
)

type CreateTipoResponse struct {
	handler.BaseResponse
	Data int `json:"data"`
}

type DeleteTipoResponse struct {
	handler.BaseResponse
	Data int `json:"data"`
}

type ShowTipoResponse struct {
	handler.BaseResponse
	Data int `json:"data"`
}

type UpdateTipoResponse struct {
	handler.BaseResponse
	Data int `json:"data"`
}

type ListTipoResponse struct {
	handler.BaseResponse
	Data []schemas.TipoResponse `json:"data"`
}
