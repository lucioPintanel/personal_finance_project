package card

import (
	"github.com/lucioPintanel/personal_finance_project/internal/handler"
	"github.com/lucioPintanel/personal_finance_project/internal/schemas"
)

type createCardResponse struct {
	handler.BaseResponse
	Data int `json:"data"`
}

type deleteCardResponse struct {
	handler.BaseResponse
	Data int `json:"data"`
}

type showCardResponse struct {
	handler.BaseResponse
	Data int `json:"data"`
}

type updateCardResponse struct {
	handler.BaseResponse
	Data int `json:"data"`
}

type listCardsResponse struct {
	handler.BaseResponse
	Data []schemas.CardResponse `json:"data"`
}
