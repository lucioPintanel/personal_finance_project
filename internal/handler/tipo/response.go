package tipo

import (
	"github.com/lucioPintanel/personal_finance_project/internal/handler"
	"github.com/lucioPintanel/personal_finance_project/internal/schemas"
)

type createTipoResponse struct {
	handler.BaseResponse
	Data int `json:"data"`
}

type deleteTipoResponse struct {
	handler.BaseResponse
	Data int `json:"data"`
}

type showTipoResponse struct {
	handler.BaseResponse
	Data int `json:"data"`
}

type updateTipoResponse struct {
	handler.BaseResponse
	Data int `json:"data"`
}

type listTipoResponse struct {
	handler.BaseResponse
	Data []schemas.TipoResponse `json:"data"`
}
