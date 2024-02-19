package transacao

import (
	"github.com/lucioPintanel/personal_finance_project/internal/handler"
	"github.com/lucioPintanel/personal_finance_project/internal/schemas"
)

type createTransacaoResponse struct {
	handler.BaseResponse
	Data int `json:"data"`
}

type deleteTransacaoResponse struct {
	handler.BaseResponse
	Data int `json:"data"`
}

type showTransacaoResponse struct {
	handler.BaseResponse
	Data int `json:"data"`
}

type updateTransacaoResponse struct {
	handler.BaseResponse
	Data int `json:"data"`
}

type listTransacaoResponse struct {
	handler.BaseResponse
	Data []schemas.TransacaoResponse `json:"data"`
}
