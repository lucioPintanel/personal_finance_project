package tipo_item

import (
	"github.com/lucioPintanel/personal_finance_project/internal/handler"
	"github.com/lucioPintanel/personal_finance_project/internal/schemas"
)

type createTipoItemResponse struct {
	handler.BaseResponse
	Data int `json:"data"`
}

type deleteTipoItemResponse struct {
	handler.BaseResponse
	Data int `json:"data"`
}

type showTipoItemResponse struct {
	handler.BaseResponse
	Data int `json:"data"`
}

type updateTipoItemResponse struct {
	handler.BaseResponse
	Data int `json:"data"`
}

type listTipoItemsResponse struct {
	handler.BaseResponse
	Data []schemas.TipoItemResponse `json:"data"`
}
