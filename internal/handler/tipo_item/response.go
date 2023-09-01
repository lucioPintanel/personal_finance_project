package tipo_item

import (
	"github.com/lucioPintanel/personal_finance_project/internal/handler"
	"github.com/lucioPintanel/personal_finance_project/internal/schemas"
)

type CreateTipoItemResponse struct {
	handler.BaseResponse
	Data int `json:"data"`
}

type DeleteTipoItemResponse struct {
	handler.BaseResponse
	Data int `json:"data"`
}

type ShowTipoItemResponse struct {
	handler.BaseResponse
	Data int `json:"data"`
}

type UpdateTipoItemResponse struct {
	handler.BaseResponse
	Data int `json:"data"`
}

type ListTipoItemsResponse struct {
	handler.BaseResponse
	Data []schemas.TipoItemResponse `json:"data"`
}
